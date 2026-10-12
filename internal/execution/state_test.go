package execution

import (
	"errors"
	"math"
	"testing"
)

func mustApply(t *testing.T, s State, typ CommandType, generation uint64) State {
	t.Helper()
	next, err := Apply(s, Command{Type: typ, ExpectedGeneration: generation})
	if err != nil {
		t.Fatalf("%s: %v", typ, err)
	}
	return next
}

func TestHappyPath(t *testing.T) {
	s := State{Status: Queued}
	s = mustApply(t, s, Claim, 0)
	if s.Generation != 1 || s.Status != Leased {
		t.Fatalf("unexpected claim: %+v", s)
	}
	s = mustApply(t, s, Start, 1)
	s = mustApply(t, s, Complete, 1)
	if s.Status != Succeeded {
		t.Fatalf("wrong terminal state: %+v", s)
	}
}

func TestStaleWorkerCannotCommitAfterReclaim(t *testing.T) {
	s := State{Status: Queued}
	s = mustApply(t, s, Claim, 0) // Worker A, generation 1
	s = mustApply(t, s, Start, 1)
	s = mustApply(t, s, ExpireSafe, 1) // repository must prove safe to retry
	s = mustApply(t, s, Claim, 1)      // Worker B, generation 2
	s = mustApply(t, s, Start, 2)

	unchanged, err := Apply(s, Command{Type: Complete, ExpectedGeneration: 1})
	if !errors.Is(err, ErrStaleGeneration) {
		t.Fatalf("expected stale-generation rejection, got %v", err)
	}
	if unchanged != s {
		t.Fatalf("failed transition mutated state: before=%+v after=%+v", s, unchanged)
	}

	s = mustApply(t, s, Complete, 2)
	if s.Status != Succeeded || s.Generation != 2 {
		t.Fatalf("current worker must succeed: %+v", s)
	}
}

func TestUnknownExternalEffectCannotBeReclaimed(t *testing.T) {
	s := State{Status: Queued}
	s = mustApply(t, s, Claim, 0)
	s = mustApply(t, s, Start, 1)
	s = mustApply(t, s, ExpireUncertain, 1)

	if s.Status != ReconciliationRequired {
		t.Fatalf("unknown effects must require reconciliation: %+v", s)
	}
	_, err := Apply(s, Command{Type: Claim, ExpectedGeneration: 1})
	if !errors.Is(err, ErrInvalidTransition) {
		t.Fatalf("ambiguous effect must not be blindly retried: %v", err)
	}
}

func TestTerminalStateIsImmutable(t *testing.T) {
	for _, terminal := range []Status{Succeeded, Failed} {
		s := State{Status: terminal, Generation: 4}
		for _, cmd := range []CommandType{Claim, Start, Complete, Fail, ExpireSafe, ExpireUncertain} {
			next, err := Apply(s, Command{Type: cmd, ExpectedGeneration: 4})
			if !errors.Is(err, ErrInvalidTransition) || next != s {
				t.Fatalf("terminal=%s cmd=%s: next=%+v err=%v", terminal, cmd, next, err)
			}
		}
	}
}

func TestClaimRejectsStaleGeneration(t *testing.T) {
	s := State{Status: Queued, Generation: 6}
	_, err := Apply(s, Command{Type: Claim, ExpectedGeneration: 5})
	if !errors.Is(err, ErrStaleGeneration) {
		t.Fatalf("wanted stale claim rejection; got %v", err)
	}
}

func TestClaimGenerationOverflow(t *testing.T) {
	s := State{Status: Queued, Generation: math.MaxUint64}
	_, err := Apply(s, Command{Type: Claim, ExpectedGeneration: math.MaxUint64})
	if !errors.Is(err, ErrGenerationOverflow) {
		t.Fatalf("wanted overflow; got %v", err)
	}
}

func TestInvalidStateOrCommand(t *testing.T) {
	_, err := Apply(State{Status: "MYSTERY"}, Command{Type: Claim})
	if !errors.Is(err, ErrInvalidState) {
		t.Fatalf("wanted invalid state; got %v", err)
	}
	_, err = Apply(State{Status: Queued}, Command{Type: "MYSTERY"})
	if !errors.Is(err, ErrInvalidTransition) {
		t.Fatalf("wanted invalid transition; got %v", err)
	}
}

func TestReconciliationRequiredRejectsAllWorkerCommands(t *testing.T) {
	state := State{
		Status:     ReconciliationRequired,
		Generation: 4,
	}

	commands := []CommandType{
		Claim,
		Start,
		Complete,
		Fail,
		ExpireSafe,
		ExpireUncertain,
	}

	for _, commandType := range commands {
		t.Run(string(commandType), func(t *testing.T) {
			next, err := Apply(state, Command{
				Type:               commandType,
				ExpectedGeneration: 4,
			})

			if !errors.Is(err, ErrInvalidTransition) {
				t.Fatalf(
					"command %s: expected invalid transition, got %v",
					commandType,
					err,
				)
			}

			if next != state {
				t.Fatalf(
					"command %s changed reconciliation state: before=%+v after=%+v",
					commandType,
					state,
					next,
				)
			}
		})
	}
}
