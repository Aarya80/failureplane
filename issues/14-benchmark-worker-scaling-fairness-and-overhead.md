# Benchmark worker scaling fairness and overhead

**Issue 14** | Depends on: #7, #8, #10, #13

**Goal:** Publish meaningful, reproducible engineering measurements.

**Implementation:** No-op, CPU WASM, mock API, skewed tenant traffic; worker counts; p50/p95/p99; saturation/recovery/fencing oracles; hardware/environment manifests.

**Acceptance:** Multiple repeat runs, raw metrics and correct comparison baselines checked in; no fabricated results.

**Evidence:** Reproducible bench command, raw JSON/CSV and analysis.
