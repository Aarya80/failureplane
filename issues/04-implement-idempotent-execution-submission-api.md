# Implement idempotent execution submission API

**Issue 04** | Depends on: #3

**Goal:** Durable creation with conflict behavior.

**Implementation:** Implement POST /v1/executions, canonical request digest, tenant-qualified transaction, initial event. Duplicate same tenant/key/request returns same ID even across response-loss retry. Same key/different digest returns 409. Enforce input size. Implement GET.

**Acceptance:** Concurrent same-key submissions create ONE row; changed digest conflicts; crash before/after commit behaves as documented; no success response prior to commit.

**Evidence:** Postgres integration tests and API contract tests.
