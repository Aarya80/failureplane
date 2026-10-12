# Implement and test execution ownership state machine

**Issue 02** | Depends on: #1

**Goal:** Establish the pure Go correctness core without network/database dependencies.

**Starting point:** Starter includes a reducer, normal/race tests and a CLI demo; inspect and extend rather than claiming they are new implementation work.

**Implementation:** Review all transitions; add table-driven invalid-transition matrix; fuzz arbitrary command sequences; define error classification and property test for monotonic generations and terminal immutability.

**Acceptance:** `go test ./...`, `go test -race ./...`, `go vet ./...` pass; stale-generation and uncertain-effect tests pass; no DB assumptions are silently implied.

**Evidence:** Test transcript and invariants mapping.
