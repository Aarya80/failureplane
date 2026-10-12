# Implement seeded worker and network fault scenarios

**Issue 13** | Depends on: #6, #11, #12

**Goal:** Reproduce failure at defined durable boundaries.

**Implementation:** Seeded DSL to kill worker at named event and inject mock provider/network failure. Controller stores scenario metadata; tests have objective invariants.

**Acceptance:** Crash-before-commit test reproducible by another engineer; ambiguous-effect test requires reconciliation; artifacts carry seed/commit.

**Evidence:** CI fault-run artifacts and sanitized execution history.
