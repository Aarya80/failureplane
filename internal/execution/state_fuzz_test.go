package execution

import "testing"

// FuzzTransitions ensures arbitrary commands cannot corrupt core invariants.
func FuzzTransitions(f *testing.F) {
	f.Add([]byte{0, 1, 2})
	f.Add([]byte{0, 1, 4, 0, 1, 2})
	f.Add([]byte{0, 1, 5, 0})
	f.Fuzz(func(t *testing.T, instructions []byte) {
		if len(instructions) > 256 {
			return
		}
		s := State{Status: Queued}
		commands := []CommandType{Claim, Start, Complete, Fail, ExpireSafe, ExpireUncertain}
		for _, b := range instructions {
			before := s
			command := Command{Type: commands[int(b)%len(commands)], ExpectedGeneration: s.Generation}
			next, err := Apply(s, command)
			if err != nil {
				if next != before {
					t.Fatalf("rejected command changed state")
				}
				continue
			}
			if next.Generation < before.Generation {
				t.Fatalf("generation decreased")
			}
			if (before.Status == Succeeded || before.Status == Failed) && next != before {
				t.Fatalf("terminal state reopened")
			}
			s = next
		}
	})
}
