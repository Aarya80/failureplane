# Worker crash runbook (future test)

Status: draft; do not execute until API, workers and leases exist.

1. Create a seeded safe-to-retry tool execution.
2. Record execution ID, worker ID, generation and journal sequence.
3. SIGKILL the active worker at a defined boundary.
4. Observe lease expiry from PostgreSQL database time.
5. Ensure a second worker claims a strictly higher generation.
6. Inject an attempted completion from the old generation; require rejection.
7. Confirm exactly one authoritative terminal state and ordered events.
8. Repeat at an external-effect boundary; require reconciliation if ambiguous.
9. Export sanitized trace, raw test oracle and environment manifest.

Never interpret an absent response as proof of absent external side effect.
