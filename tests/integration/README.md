# Integration test suite (planned)

Milestones 3-6 will add real PostgreSQL integration tests, simultaneous
claimers, idempotency races, lease expiry tests, journal atomicity, and
fault-injected external-effect reconciliation. No mock should replace Postgres
for transaction-correctness evidence.
