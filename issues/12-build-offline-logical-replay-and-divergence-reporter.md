# Build offline logical replay and divergence reporter

**Issue 12** | Depends on: #9, #11

**Goal:** Replay controlled host effects without live external side effects.

**Implementation:** Pin tool/input digests, sequence recorded effect responses, reject unrecorded capabilities, compare normalized outputs/critical events, INCOMPLETE on missing evidence.

**Acceptance:** Identical replay passes; intentionally modified effect response diverges; live network is unavailable in replay.

**Evidence:** Committed golden replay fixtures and CLI result.
