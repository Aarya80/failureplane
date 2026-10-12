-- Local development rollback. DESTRUCTIVE: removes all stored executions.
BEGIN;
DROP TABLE IF EXISTS effects;
DROP TABLE IF EXISTS execution_events;
DROP TABLE IF EXISTS attempts;
DROP TABLE IF EXISTS executions;
DROP TABLE IF EXISTS tenants;
COMMIT;
