# Current project status

Last packaged: 2026-10-11

VERIFIED LOCALLY:
- `go test ./...` passes.
- `go test -race ./...` passes.
- `go vet ./...` passes.
- `go run ./cmd/failureplane demo` demonstrates stale-generation rejection.
- Pure state machine models uncertain-effect -> reconciliation-required.
- Python stdlib loopback experiment proves a provider can accept a write even
  when the caller receives no response, and shows provider idempotency.

DESIGNED, NOT IMPLEMENTED:
- PostgreSQL schema (SQL provided; database migration not applied here).
- HTTP API, scheduler, worker, Rust/Wasmtime runner, effect broker, replay,
  real fault controller, SDK, dashboard, benchmarks, deployment.

NEXT TASK:
- Issue #1: review invariant/ADR, then issue #3 persistence schema harness or
  issue #2 state-machine expansion. The first five issues are in `issues/`.

IMPORTANT:
The `ExpireSafe` command requires the future durable repository to prove
lease expiry and safe retry using authoritative DB/effect state. The in-memory
demo is not evidence of distributed recovery.

## Issue #1 local verification — 2026-10-11

VERIFIED:
- TestReconciliationRequiredRejectsAllWorkerCommands passes
  for all six existing worker commands.
- go test ./... passed after the new regression test.
- go test -race ./... passed after the new regression test.
- go vet ./... passed after the new regression test.
- git diff --check passed after adding failure timelines.
- Stale-owner and ambiguous-effect timelines documented.

BOUNDARIES:
- Reconciliation enforcement is currently in-memory only.
- PostgreSQL lease fencing and effect brokerage are planned.
- GitHub CI passed on initial commit 92eaeed; Issue #1 changes
  have not yet been pushed or verified by CI.

NEXT:
- Commit Issue #1 changes and verify GitHub CI.
- Begin Issue #2: strengthen state-machine property and fuzz tests.
