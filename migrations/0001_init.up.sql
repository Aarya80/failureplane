-- PLANNED persistence milestone. PostgreSQL 16+.
-- Not applied or tested by the Go in-memory starter.
-- SECURITY WARNING: no RLS is enabled yet. Keep this database local-only.
BEGIN;

CREATE TABLE tenants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE executions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id),
    tool_digest TEXT NOT NULL,
    input_digest TEXT NOT NULL,
    request_digest TEXT NOT NULL,
    idempotency_key TEXT,
    status TEXT NOT NULL DEFAULT 'QUEUED'
        CHECK (status IN ('QUEUED', 'LEASED', 'RUNNING', 'RETRY_WAIT',
                         'SUCCEEDED', 'FAILED', 'CANCELED',
                         'RECONCILIATION_REQUIRED')),
    generation BIGINT NOT NULL DEFAULT 0 CHECK (generation >= 0),
    lease_owner TEXT,
    lease_expires_at TIMESTAMPTZ,
    next_run_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    attempt_count INT NOT NULL DEFAULT 0 CHECK (attempt_count >= 0),
    max_attempts INT NOT NULL DEFAULT 3 CHECK (max_attempts > 0),
    next_event_seq BIGINT NOT NULL DEFAULT 1 CHECK (next_event_seq > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CHECK ((lease_owner IS NULL AND lease_expires_at IS NULL)
        OR (lease_owner IS NOT NULL AND lease_expires_at IS NOT NULL))
);

CREATE UNIQUE INDEX executions_idempotency_uq
    ON executions (tenant_id, idempotency_key)
    WHERE idempotency_key IS NOT NULL;
CREATE INDEX executions_tenant_created_idx
    ON executions (tenant_id, created_at DESC);
CREATE INDEX executions_ready_idx
    ON executions (next_run_at, created_at)
    WHERE status IN ('QUEUED', 'RETRY_WAIT');
CREATE INDEX executions_expired_lease_idx
    ON executions (lease_expires_at)
    WHERE status IN ('LEASED', 'RUNNING');

CREATE TABLE attempts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    execution_id UUID NOT NULL REFERENCES executions(id),
    generation BIGINT NOT NULL CHECK (generation > 0),
    worker_id TEXT NOT NULL,
    started_at TIMESTAMPTZ,
    finished_at TIMESTAMPTZ,
    outcome TEXT,
    UNIQUE (execution_id, generation)
);
CREATE INDEX attempts_execution_idx
    ON attempts (execution_id, generation DESC);

CREATE TABLE execution_events (
    execution_id UUID NOT NULL REFERENCES executions(id),
    seq BIGINT NOT NULL CHECK (seq > 0),
    event_type TEXT NOT NULL,
    attempt_generation BIGINT,
    payload JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (execution_id, seq)
);

CREATE TABLE effects (
    execution_id UUID NOT NULL REFERENCES executions(id),
    effect_key TEXT NOT NULL,
    request_digest TEXT NOT NULL,
    external_idempotency_key TEXT,
    state TEXT NOT NULL CHECK (state IN ('INTENT_RECORDED', 'DISPATCHING',
        'SUCCEEDED', 'FAILED_FINAL', 'UNKNOWN')),
    response_digest TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (execution_id, effect_key)
);

COMMIT;
