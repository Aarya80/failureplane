# Implement initial PostgreSQL migrations and test harness

**Issue 03** | Depends on: #1, #2

**Goal:** Create a real, transaction-capable persistence foundation.

**Starting point:** `migrations/0001_init.up.sql` is a design draft, not applied/validated.

**Implementation:** Pin DB version; write reproducible migration apply/rollback tool; validate schema/indexes/constraints using real Postgres in Docker; add transaction rollback test for state + event append. Decide row lock sequence allocation.

**Acceptance:** Fresh `docker compose` + migration works from clean environment, rollback works on disposable DB, constraint tests pass, rollback atomicity confirmed.

**Evidence:** CI integration test and Postgres command output.
