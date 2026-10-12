// failureplane is currently a demo CLI for the in-memory state model.
// It is NOT a functioning durable executor or sandbox.
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Aarya80/failureplane/internal/execution"
)

func usage() {
	fmt.Println("FailurePlane — engineering starter (not production-ready)")
	fmt.Println("Usage: failureplane demo | help")
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "help" {
		usage()
		return
	}
	if os.Args[1] != "demo" {
		usage()
		os.Exit(2)
	}
	if err := demo(); err != nil {
		fmt.Fprintln(os.Stderr, "demo failed:", err)
		os.Exit(1)
	}
}

func demo() error {
	s := execution.State{Status: execution.Queued}
	fmt.Println("In-memory ownership demonstration (no database/network):")
	fmt.Printf("Initial: %s (generation %d)\n", s.Status, s.Generation)

	steps := []execution.Command{
		{Type: execution.Claim, ExpectedGeneration: 0},
		{Type: execution.Start, ExpectedGeneration: 1},
		{Type: execution.ExpireSafe, ExpectedGeneration: 1},
		{Type: execution.Claim, ExpectedGeneration: 1},
		{Type: execution.Start, ExpectedGeneration: 2},
	}
	for _, cmd := range steps {
		next, err := execution.Apply(s, cmd)
		if err != nil {
			return err
		}
		s = next
		fmt.Printf("%-14s -> %s (generation %d)\n", cmd.Type, s.Status, s.Generation)
	}

	_, err := execution.Apply(s, execution.Command{Type: execution.Complete, ExpectedGeneration: 1})
	if !errors.Is(err, execution.ErrStaleGeneration) {
		return errors.New("old worker was not rejected")
	}
	fmt.Println("Old worker completion -> REJECTED (stale generation)")
	s, err = execution.Apply(s, execution.Command{Type: execution.Complete, ExpectedGeneration: 2})
	if err != nil {
		return err
	}
	fmt.Printf("New worker completion -> %s (generation %d)\n", s.Status, s.Generation)
	fmt.Println("NEXT: enforce the same invariant in a PostgreSQL transaction.")
	return nil
}
