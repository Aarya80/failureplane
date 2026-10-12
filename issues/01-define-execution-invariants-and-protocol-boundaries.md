# Define execution invariants and protocol boundaries

**Issue 01** | Depends on: none

**Goal:** Freeze the smallest correct system contract before new infrastructure is added.

**Implementation:** Document execution lifecycle, generation monotonicity, lease authority, allowed state transitions, durability acceptance point, event ordering, retry classifications, external-effect ambiguity and replay limits. Put ADR links in the README.

**Acceptance:** Run and explain the real loopback ambiguous-effect proof `python3 examples/faults/ambiguous_effect_spike.py`; architecture/guarantee docs accurately match current code; one failure timeline includes a stale worker and one includes an ambiguous external write; explicitly list non-guarantees.

**Evidence:** ADR diff, reviewed invariant checklist, linked reducer tests.
