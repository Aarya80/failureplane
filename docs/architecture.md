# Architecture (intended, not yet built)

## Purpose
A developer runs a bounded tool or async workload and can inspect its durable
state, crash recovery, effect outcomes and logical replay evidence.

```text
CLI / SDK -> Go API -> PostgreSQL (durable authority)
                          ^  |
                          |  v
                   Go scheduler / worker(s)
                          |
                          v
                   Rust Wasmtime host
                          |
                   capability imports
                          |
                          v
                   Go effect broker -> allowlisted external services

Fault controller -> workers / network + replay reader -> event history
OTel instrumentation -> traces/metrics -> inspector UI
```

## Responsibilities
1. API: authenticate, authorize, validate input, reserve idempotency key,
   transactionally create execution and first event, return ID after commit.
2. Repository: transactional state and event changes, owned generation, lease
   time using DB clock, idempotent submit and effect states.
3. Scheduler: claim due tasks, cap concurrency, fair dispatch and reclaims.
4. Worker: run attempt, heartbeat, handle cancellation and result commit.
5. Runner: isolated Wasmtime component invocation with approved imports.
6. Effect broker: allowed host calls, record intent and outcomes, reconcile
   ambiguity and reject unauthorized destinations.
7. Replay: read immutable tool/input + recorded host boundary, offline compare.
8. Developer UX: CLI/TS SDK/inspector, initially read-only/allowlisted demo.

## Claim transaction (conceptual pseudocode)
1. BEGIN
2. SELECT id FROM executions WHERE status='QUEUED' AND next_run_at <= now()
   ORDER BY next_run_at, created_at FOR UPDATE SKIP LOCKED LIMIT 1
3. UPDATE selected row: status=LEASED, generation=generation+1,
   lease_owner=worker, lease_expires_at=DB_clock + duration;
   append corresponding event; create attempt (same transaction)
4. COMMIT; return claimed work

## Fenced result commit (conceptual pseudocode)
- BEGIN
- UPDATE executions ... WHERE id=? AND generation=? AND
  lease_owner=? AND lease_expires_at > clock_timestamp() AND status='RUNNING'
- Require exactly one updated row; append journal event and mark attempt,
  allocating per-execution sequence under row lock
- COMMIT

## Major correctness boundaries
- Transaction commits, not API responses, define durable acceptance.
- Lease expiry permits reacquisition only after effect-safety check.
- Generation fencing prevents old owners from writing terminal state.
- Side effects are not exactly-once without downstream coordination.
- WASM sandbox is restricted but not a complete isolation proof by itself.
- Logical replay uses recorded host effects; it is not process time travel.

## Scaling path
One PostgreSQL primary first; workers scale horizontally while claims are
controlled in DB. Introduce cache/object store/remote broker only when
instrumentation justifies them. No Kubernetes or leader election by default.
