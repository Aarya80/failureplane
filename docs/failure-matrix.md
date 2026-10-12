# Fault and recovery matrix (planned acceptance suite)

| Injection | Expected outcome | Correctness oracle |
|---|---|---|
| API dies before DB commit | Not accepted; retry safe | Unique tenant/key record, no partial event |
| API dies after DB commit | Repeated submission returns same ID | No duplicate execution |
| Worker dies before guest start | Lease expiry permits reclaim | Generation increments and history explains it |
| Worker dies during pure guest computation | Safe additional attempt | One authoritative terminal commit |
| Worker dies after possible external effect | STOP or reconcile uncertainty | No unguarded repeat of non-idempotent write |
| Old worker returns after lease transfer | Stale result rejected | No state/event mutation |
| Two workers claim same task | Only one generation's owner | One claim succeeds for that generation |
| Database unavailable | Stop authoritative progress | No fabricated success acknowledgement |
| Network timeouts | Classify unknown vs known-safe | Effect record + correct disposition |
| Cache unavailable | Immutable artifact miss/fallback | Correctness preserved |
| Queue overloaded | Admission caps and bounded memory | No unbounded growth, status 429/503 |
| Malformed request/component | Reject before execution | No host panic; response categorized |
| Guest infinite loop | Trap / wall-time/fuel limit | Host remains responsive |
| Guest memory abuse | Bounded memory/failure | Host resource metrics within budget |
| Disallowed host call | Capability denied | No file/network access outside policy |
| Fault controller restart | Reconstruct from durable test record | Same logical injection point/seed |
| Trace exporter unavailable | Run continues with telemetry-loss metric | Bounded exporter buffers |

All test outcomes are future requirements; none are claimed achieved by starter.
