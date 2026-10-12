# Current project status

Snapshot: 2026-10-11

## Current milestone

Phase 0 proof completed.
Issue #1 complete.
Issue #2 complete: local tests and GitHub CI verified.
Issues #3–15 planned.

Locally verified Issue #2 test commits: 3e4d496 and d297b74.
GitHub Actions run 38198766481 passed on commit 0651280.

## Verified — Issue #1

- Execution invariants and failure timelines documented.
- Pure reducer rejects stale-generation completion.
- Reconciliation-required state rejects ordinary worker commands.
- Python mock-provider demonstrates an accepted write with lost response.
- Provider-idempotent retry demonstrated in the mock experiment.
- GitHub Actions run 38193694278 completed successfully.

## Verified — Issue #2

- Existing pure Go ownership reducer reviewed.
- Existing arbitrary command-sequence fuzz test retained.
- Stale-generation fuzz test added in commit 3e4d496.
- Invalid-transition matrix added in commit d297b74.
- All 27 forbidden state-command combinations passed.
- Stale-generation fuzz seeds passed.
- Active stale-generation fuzz run completed 56,205 executions
  without a reported failure.
- go test ./... passed.
- go test -race ./... passed.
- go vet ./... produced no diagnostics.
- gofmt and git diff --check produced no issues.

The fuzz execution count is test evidence, not a
performance benchmark.

## Issue #2 invariant-to-test mapping

| Property | Tests |
| --- | --- |
| Generation never decreases | FuzzTransitions |
| Stale generation cannot mutate state | TestStaleWorkerCannotCommitAfterReclaim, FuzzStaleGenerationCannotMutate |
| Terminal states cannot reopen | TestTerminalStateIsImmutable, FuzzTransitions |
| Invalid transitions preserve state | TestInvalidTransitionMatrix |
| Unknown effect prevents ordinary retry | TestUnknownExternalEffectCannotBeReclaimed, TestReconciliationRequiredRejectsAllWorkerCommands |
| Generation overflow rejected | TestClaimGenerationOverflow |
| Invalid state and command rejected | TestInvalidStateOrCommand |

Error classification:
- ErrInvalidState: unrecognized execution status.
- ErrStaleGeneration: expected generation mismatch.
- ErrInvalidTransition: forbidden state-command transition.
- ErrGenerationOverflow: claim would overflow generation counter.

## Boundaries and non-guarantees

The pure reducer does not enforce worker identity,
lease expiry, database transactions, or durable effects.

ExpireSafe requires the future repository to establish
lease expiry and prove retry safety before invocation.

No exactly-once execution or arbitrary external-effect
guarantee is provided.

PostgreSQL durability, real workers, Wasmtime isolation,
external-effect brokerage, logical replay, and benchmarks
are not implemented.

## Next actions

1. Commit and push the final Issue #2 status update.
2. Verify CI on the documentation-only commit.
3. Begin Issue #3: PostgreSQL migrations and integration harness.
