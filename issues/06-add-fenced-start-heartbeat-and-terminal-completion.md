# Add fenced start heartbeat and terminal completion

**Issue 06** | Depends on: #5

**Goal:** Complete the ownership protocol safely.

**Implementation:** State transition SQL must check generation, owner, status and DB lease time; insert journal events transactionally. Heartbeats extend only current leases. Handle lost DB connectivity without guessed success.

**Acceptance:** Replay delayed completion from previous generation and show zero changed rows/events. Simulate DB outage.

**Evidence:** Integration suite and recovery trace.
