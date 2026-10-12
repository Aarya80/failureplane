# Define and implement worker-to-runner boundary

**Issue 10** | Depends on: #7, #9

**Goal:** Run isolated guest from worker with explicit lifecycle and structured outcomes.

**Implementation:** Versioned request/response envelope, process supervision, timeout/cancel behavior, bounded stdin/stdout or IPC; no arbitrary host secret exposure.

**Acceptance:** SIGKILL runner causes classified attempt failure; malformed output rejected; worker remains alive.

**Evidence:** E2E worker + runner trace.
