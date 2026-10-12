# Maturity gates, not arbitrary deadlines

## Phase 0 — Risk proof
Deliver fake provider with accept-and-drop, a persisted effect-intent design,
WASM import-denial spike. Exit: explicit ambiguous-effect demonstration.

## Phase 1 — Small real executor
API + DB + one worker + one deterministic WIT tool. Exit: a durable accepted
execution and readable terminal result after restarting services.

## Phase 2 — Correctness
Atomic journal, idempotency, claims, lease fencing, fail classifications and
real Postgres concurrency tests. Exit: invariants pass under injected crashes.

## Phase 3 — Multiple workers
Worker discovery, ownership transfer, bounded concurrency and fairness tests.
Exit: correctness preserved with several concurrent workers.

## Phase 4 — Effects + replay
Brokered effects, UNKNOWN state, offline logical replay, seeded fault DSL.
Exit: reproducible ambiguous-effect incident without duplicate writes.

## Phase 5 — Operational platform
Auth, tenant isolation, sandbox limits, OTEL, Docker, release automation.
Exit: staging reliability/security gate, documented threat model.

## Phase 6 — Performance
Profile, optimize and publish reproducible raw benchmarks and tradeoffs.
Exit: benchmark scripts and results reproduce without personal guidance.

## Phase 7 — Developer experience / launch
CLI, TypeScript SDK, inspector, sample workloads, easy self-hosting.
Exit: stranger runs failure demo following README.

## Phase 8 — External feedback
Real developer integrations, issue/release history, bug fixes.
Exit: credible evidence someone else used a meaningful workflow.

## Portfolio quality gates
- GitHub-ready: code/tests/docs from real commit and honest current status.
- Portfolio-ready: end-to-end live demo and clear technically original work.
- Resume-ready: durable multi-worker recovery + sandbox + effects/replay +
  reproducible tests/benchmark. Do not use planned numbers as achievements.
- Interview-ready: defend code, race/failure decisions and evidence.
- Public-ready: safe deploy, documented limits and external-use testing.
