// Package execution defines the pure execution-ownership state machine.
// It performs NO I/O. Database transactions and lease expiry checks are
// deliberately out of scope until the persistence milestone.
package execution

import (
	"errors"
	"math"
)

type Status string

const (
	Queued                 Status = "QUEUED"
	Leased                 Status = "LEASED"
	Running                Status = "RUNNING"
	Succeeded              Status = "SUCCEEDED"
	Failed                 Status = "FAILED"
	ReconciliationRequired Status = "RECONCILIATION_REQUIRED"
)

type CommandType string

const (
	Claim           CommandType = "CLAIM"
	Start           CommandType = "START"
	Complete        CommandType = "COMPLETE"
	Fail            CommandType = "FAIL"
	ExpireSafe      CommandType = "EXPIRE_SAFE"
	ExpireUncertain CommandType = "EXPIRE_UNCERTAIN"
)

var (
	ErrStaleGeneration    = errors.New("stale execution generation")
	ErrInvalidTransition  = errors.New("invalid execution transition")
	ErrGenerationOverflow = errors.New("execution generation overflow")
	ErrInvalidState       = errors.New("invalid execution state")
)

type State struct {
	Status     Status
	Generation uint64
}

type Command struct {
	Type               CommandType
	ExpectedGeneration uint64
}

// Apply validates an ownership-sensitive state transition.
//
// The caller MUST verify actual lease expiry against database time before
// issuing ExpireSafe / ExpireUncertain. The caller MUST inspect durable
// effect records to determine whether re-execution is safe. This reducer
// never infers effect safety from a worker timeout.
//
// The future PostgreSQL repository MUST update the current state and append
// an event atomically, conditional on generation + lease ownership.
func Apply(state State, command Command) (State, error) {
	switch state.Status {
	case Queued, Leased, Running, Succeeded, Failed, ReconciliationRequired:
	default:
		return state, ErrInvalidState
	}

	if command.ExpectedGeneration != state.Generation {
		return state, ErrStaleGeneration
	}

	next := state
	switch command.Type {
	case Claim:
		if state.Status != Queued {
			return state, ErrInvalidTransition
		}
		if state.Generation == math.MaxUint64 {
			return state, ErrGenerationOverflow
		}
		next.Generation++
		next.Status = Leased
	case Start:
		if state.Status != Leased {
			return state, ErrInvalidTransition
		}
		next.Status = Running
	case Complete:
		if state.Status != Running {
			return state, ErrInvalidTransition
		}
		next.Status = Succeeded
	case Fail:
		if state.Status != Running && state.Status != Leased {
			return state, ErrInvalidTransition
		}
		next.Status = Failed
	case ExpireSafe:
		if state.Status != Running && state.Status != Leased {
			return state, ErrInvalidTransition
		}
		next.Status = Queued
	case ExpireUncertain:
		if state.Status != Running && state.Status != Leased {
			return state, ErrInvalidTransition
		}
		next.Status = ReconciliationRequired
	default:
		return state, ErrInvalidTransition
	}
	return next, nil
}
