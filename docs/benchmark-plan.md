# Reproducible benchmark plan

## Core comparisons
A. In-process no-op transition baseline.
B. Single-worker PostgreSQL state/journal baseline.
C. Multi-worker + runner + broker + journal full-path baseline.

## Workloads
No-op; CPU-bound WASM; small/large JSON; stubbed external call; fake provider
accept/drop; worker-kill churn; burst over capacity; tenant traffic skew.

## Fixed configuration per run
Git commit, module checksum, Docker image digest, Go/Rust/Wasmtime/Postgres
versions, OS/kernel, CPU, RAM, disk, DB pool, worker count, input seed,
workload distribution, cold/warm cache, failure script.

## Metrics
Ops/s, p50/p95/p99 by stage, CPU/RSS, DB lock waits, bytes/event,
cache hit, queue depth, fault recovery latency, stale commit rejections,
unknown-effect counts, replay divergence, correctness oracle results.

## Protocol
1. Verify correctness baseline before optimizing.
2. Warm up and separately label cold-start runs.
3. Run repeated independent trials; save raw samples.
4. Present distribution and sample size, not cherry-picked best trial.
5. Publish scripts, hardware specs, evidence and known confounders.
6. Benchmark against comparable baselines with equivalent guarantees.

All numbers are MEASUREMENT TARGETS, not project accomplishments.
