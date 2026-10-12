# Implement brokered effects and reconciliation state

**Issue 11** | Depends on: #8, #10

**Goal:** Resolve the central real-world failure ambiguity honestly.

**Implementation:** Persist effect intent before dispatch; stable per-effect key/request digest; synthetic provider accept/drop; states SUCCEEDED, FAILED_FINAL, UNKNOWN; safe adapter reconciliation/stop.

**Acceptance:** Simulated downstream write accepted then response lost never triggers blind second write without allowed policy.

**Evidence:** Fault experiment with mock provider write ledger.
