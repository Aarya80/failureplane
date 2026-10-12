# Benchmark plan (not executed)

Benchmarks are NOT implemented in this starter. Proposed first comparisons:
1. Pure in-process no-op state changes (low-level reducer baseline).
2. Postgres-backed queued/claimed no-op execution, one worker.
3. Full journaled executor with four workers and forced crash/recovery.

Record: git commit, Go/Rust/Postgres versions, CPU/RAM/OS, Docker settings,
workload seed, concurrency, tool digest, raw latency samples, p50/p95/p99,
throughput, database contention, lease recovery time and correctness oracles.

Do not publish synthetic placeholder metrics as empirical results.
