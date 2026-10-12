# Prototype restricted Rust Wasmtime tool host

**Issue 09** | Depends on: #1

**Goal:** Prove actual bounded guest execution with WIT, no unapproved host access.

**Implementation:** Pin audited Wasmtime version, implement one deterministic WIT `run` function, host import allowlist, input/output limits, fuel/epoch/wall time and memory caps. Run sandbox host with least privilege.

**Acceptance:** Valid deterministic guest succeeds; forbidden imports rejected; infinite loop and memory hog fail without host crash; exact runtime version documented.

**Evidence:** Security-adversarial test logs, minimal runnable Rust crate.
