# ADR-0001 — PostgreSQL as the initial durable authority

Status: **Accepted for design, implementation pending**
Date: 2026-10-11

## Context
The product needs durable acceptance, atomic execution state changes, event
history, idempotent submission, and safe ownership changes between workers.

## Decision
Start with a **single PostgreSQL primary** as the source of truth for
executions, attempts, lease ownership, generation numbers, and ordered events.
Use `FOR UPDATE SKIP LOCKED` only for work-claim selection. A claim updates the
execution row, allocates a new generation, creates an attempt, and appends the
journal event in the same transaction. The same applies to terminal commits.

Workers DO NOT decide that their lease is live using local wall clocks.
The repository must use database time in compare-and-swap predicates.

## Alternatives rejected for version 0
- Redis/BullMQ for scheduling: queue delivery and relational state would need
  extra coordination to make ownership/events atomic.
- Kafka as source of truth: strong fit for event throughput, unnecessary
  machinery for the initial scope and complicated claim semantics.
- etcd/Consul leases: premature additional coordination service.
- Temporal: mature alternative, but hides the mechanisms that are the project's
  explicit engineering-learning and failure-testing purpose.

## Guarantees
- Commit to database before reporting durable acceptance.
- Monotonically increasing ownership generation.
- Stale-generation result commits rejected by conditional SQL updates.
- No claim of exactly-once arbitrary external effects.

## Consequences and risks
- Database contention may limit throughput; profile before replacing it.
- A DB outage pauses new authoritative transitions; workers must not assume a
  successful completion was committed.
- Until row-level security is implemented, DB access is development-only.

## Verification
Concurrent two-worker claim race; stale completion after reclaim; crash before
and after transaction commit; database outage/restart; journal atomicity.
