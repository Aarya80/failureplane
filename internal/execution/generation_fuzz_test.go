package execution

import (
	"errors"
	"testing"
)

// FuzzStaleGenerationCannotMutate verifies that commands from older
// generations cannot change an execution's current state.
//
// This tests only the pure reducer. Lease expiry and effect safety
// must be established by the future durable repository.
func FuzzStaleGenerationCannotMutate(f *testing.F) {
	// Cover every reachable status with both early and later claims.
	for phase := uint8(0); phase < 6; phase++ {
		f.Add(uint8(0), phase, uint8(0))
		f.Add(uint8(7), phase, uint8(3))
	}

	commands := []CommandType{
		Claim,
		Start,
		Complete,
		Fail,
		ExpireSafe,
		ExpireUncertain,
	}

	f.Fuzz(func(t *testing.T, recoveries, phase, oldGeneration uint8) {
		state := State{Status: Queued}

		// Construct a reachable state with generation >= 1.
		// Bound the work so fuzz inputs cannot cause huge loops.
		claims := 1 + int(recoveries%16)

		for i := 0; i < claims; i++ {
			state = mustApply(t, state, Claim, state.Generation)

			if i+1 < claims {
				state = mustApply(t, state, ExpireSafe, state.Generation)
			}
		}

		// Exercise every status reachable through ordinary commands.
		switch phase % 6 {
		case 0: // LEASED
		case 1: // RUNNING
			state = mustApply(t, state, Start, state.Generation)
		case 2: // QUEUED after safe expiry
			state = mustApply(t, state, ExpireSafe, state.Generation)
		case 3: // SUCCEEDED
			state = mustApply(t, state, Start, state.Generation)
			state = mustApply(t, state, Complete, state.Generation)
		case 4: // FAILED
			state = mustApply(t, state, Fail, state.Generation)
		case 5: // RECONCILIATION_REQUIRED
			state = mustApply(t, state, ExpireUncertain, state.Generation)
		}

		// Always select a generation strictly older than the current one.
		staleGeneration := uint64(oldGeneration) % state.Generation

		for _, commandType := range commands {
			next, err := Apply(state, Command{
				Type:               commandType,
				ExpectedGeneration: staleGeneration,
			})

			if !errors.Is(err, ErrStaleGeneration) {
				t.Fatalf(
					"status=%s current=%d stale=%d command=%s: got %v",
					state.Status,
					state.Generation,
					staleGeneration,
					commandType,
					err,
				)
			}

			if next != state {
				t.Fatalf(
					"stale command mutated state: before=%+v after=%+v",
					state,
					next,
				)
			}
		}
	})
}
