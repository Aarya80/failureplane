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
