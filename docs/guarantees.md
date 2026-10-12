# Guarantees and non-guarantees

**Current starter (VERIFIED):** Go pure state reducer rejects stale generation,
models safe versus uncertain expiration, rejects invalid/terminal transitions.
Tests pass under `go test` and `go test -race`. THIS DOES NOT YET PROVIDE
DURABLE ACCEPTANCE, NETWORKED WORKERS, A SANDBOX OR REPLAY.

**Planned after implementation and verification:**
1. Durable acceptance only after PostgreSQL commit; the same request +
   tenant/idempotency key returns the original result; mismatched digest = 409.
2. Per-execution ordered append-only event history in one transaction with
   state changes.
3. Generation fencing on worker result writes; DB is authoritative for lease
   expiry and current ownership.
4. At-least-once **attempt execution** under retry conditions.
5. Reconciliation or stop for unknown external effect outcomes.
6. Offline logical replay when immutable tool, input and host effects are
   fully captured; otherwise INCOMPLETE rather than a guessed result.

**Never claim:**
- Exactly-once arbitrary external API effects.
- That a worker crash means the previous operation did not occur.
- That Wasmtime means arbitrary hostile guests are automatically safe.
- That a network timeout proves execution failed.
- That replay reproduces native scheduling, physical time or external writes.
- Production-ready isolation before threat modeling and adversarial tests.

## First invariant
If a worker holds generation 1 and ownership transfers to generation 2, a
completion bearing generation 1 must not change the execution's current state.
First unit test: `TestStaleWorkerCannotCommitAfterReclaim`.

## Database test gate
This invariant must eventually be repeated against live PostgreSQL with two
simultaneous connections and deliberate transaction delays.

## Failure timelines (Issue #1)

These timelines distinguish behavior demonstrated by the pure
reducer from behavior that requires future durable implementation.

### Timeline A: Stale worker after lease expiry

1. Worker A claims an execution at generation 1 and starts.
2. Worker A becomes unresponsive.
3. The future PostgreSQL repository verifies lease expiry using
   database time and checks durable effects before allowing retry.
4. If retry is safe, the execution returns to QUEUED.
5. Worker B claims the execution at generation 2.
6. Worker A attempts to commit completion at generation 1.
7. The stale completion is rejected without changing current state.
8. Worker B can complete while holding valid ownership.

VERIFIED: The pure reducer rejects the stale generation.

PLANNED: Transactional lease validation, owner verification,
durable effect checks, and fenced PostgreSQL commits.

### Timeline B: External write with a lost response

1. A worker requests an authorized external write.
2. The provider accepts the write.
3. The provider response is lost.
4. The caller observes an UNKNOWN outcome.
5. The execution must not be blindly re-executed.
6. If the outcome cannot be safely established, the execution
   requires reconciliation.
7. Re-execution requires evidence of safety, such as verified
   provider idempotency or status-query reconciliation.

VERIFIED: The local Python mock-provider spike demonstrates an
accepted write, lost response, and provider-idempotent retry.

VERIFIED: The pure reducer rejects ordinary worker commands
from RECONCILIATION_REQUIRED.

PLANNED: Durable effect intent records, effect brokerage,
reconciliation workflows, and database-enforced retry policy.

### Explicit limitations

- Neither timeline demonstrates real distributed crash recovery.
- Generation fencing cannot undo external operations.
- Provider idempotency depends on the provider's actual guarantees.
- A timeout does not prove that an operation failed.
- Logical replay cannot recreate arbitrary external side effects.
