# FailurePlane

**A proposed durable, sandboxed execution plane for reproducible tool failures.**

> **Status: engineering starter, NOT a deployed execution platform.**
> The code currently includes a pure Go state machine, unit/fuzz tests, and
> an in-memory worker-ownership demonstration. PostgreSQL, worker processes,
> sandbox execution, fault injection, replay, the SDK and the dashboard are
> future milestones. See [current status](docs/STATUS.md).

## Why build it?

A tool may complete externally and crash before persisting its result; a worker
can come back after its lease expires; a timeout can leave an effect ambiguous.
FailurePlane will make these failure boundaries inspectable and testable rather
than hiding them behind blind retries.

## First verified demo

Requires Go 1.23+ (or current supported Go version).

```bash
go test ./...
go test -race ./...
go run ./cmd/failureplane demo
python3 examples/faults/ambiguous_effect_spike.py
```

The Python experiment separately demonstrates an external write that succeeds
despite its response being dropped. It requires provider idempotency for safe
retries and does not simulate a complete distributed executor.

The Go demo simulates claim -> start -> safe lease expiration -> reclaim by a
new generation -> rejection of an old worker's completion. It is **in-memory**:
it does not imply a database-backed lease, multi-process worker or durable
execution guarantee yet.

## Core intended capabilities

- Durable acceptance and transactional event history in PostgreSQL.
- Workers with expiring leases and fenced completion writes.
- Rust/Wasmtime constrained execution for explicitly approved tool modules.
- Brokered external effects with an explicit UNKNOWN/reconciliation state.
- Offline logical replay using pinned inputs/tools and recorded host effects.
- Seeded crash and network fault experiments with objective correctness tests.
- Practical Go CLI / TypeScript SDK / focused execution inspector.

## Non-guarantees

Do not claim exactly-once arbitrary external API side effects. A downstream
write may succeed before the worker learns the outcome. A timeout or worker
crash does not prove that an effect did not happen.

Logical replay does not recreate physical time, arbitrary thread schedules,
external production writes or unrecorded nondeterminism.

## Planned architecture

```text
Client / CLI / SDK
        |
        v
 Go control plane <----> PostgreSQL (execution authority / journal)
        |
        v
 Scheduler + Go worker(s)
        |
        v
 Rust / Wasmtime component runner
        |
        v
 Approved-effect broker -> allowlisted external APIs

 Fault controller / offline replay <-- persisted evidence
```

See [architecture](docs/architecture.md), [guarantees](docs/guarantees.md),
[security model](docs/security-model.md), and
[ADRs](docs/adr/).

## Planned database environment

```bash
docker compose up -d postgres
# PostgreSQL-only development service; no API server is implemented yet.
```

`migrations/0001_init.up.sql` is a planned initial schema, not applied or
verified by this starter. Do not run it against a real/production database.

## Roadmap / issues

Read [roadmap](docs/roadmap.md) and the 15 issues in [`issues/`](issues/).
Start with design invariants and a real Postgres-backed fencing test; do not
build the dashboard or add AI providers before correctness is established.

## Benchmarks

None have been performed. See [benchmark methodology](docs/benchmark-plan.md).
Never use hypothetical throughput, user counts or correctness percentages as
verified resume accomplishments.

## Contributing / security / license

See [CONTRIBUTING.md](CONTRIBUTING.md), [SECURITY.md](SECURITY.md) and [LICENSE](LICENSE).
