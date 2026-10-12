# Make ordered event journal atomic

**Issue 08** | Depends on: #3, #5, #6

**Goal:** Reconstruct and inspect execution transitions.

**Implementation:** Allocate monotonically increasing per-execution sequence under row lock. Append event in same transaction as execution update. Reject attempts to infer order from wall clock.

**Acceptance:** Transaction failure leaves neither transition nor event; simultaneous changes have unambiguous ordered events.

**Evidence:** Parallel SQL tests and event trace.
