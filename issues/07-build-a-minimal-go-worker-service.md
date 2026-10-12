# Build a minimal Go worker service

**Issue 07** | Depends on: #5, #6

**Goal:** First real task claim/execute/commit loop.

**Implementation:** Worker identity, bounded polling, heartbeats, graceful shutdown, cancellation-safe context, backoff and noop tool adapter initially (NOT arbitrary guest code).

**Acceptance:** Restart worker; durable task completes; no result committed from stale worker; bounded resources.

**Evidence:** Local two-process demo.
