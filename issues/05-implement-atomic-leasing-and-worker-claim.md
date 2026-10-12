# Implement atomic leasing and worker claim

**Issue 05** | Depends on: #3, #4

**Goal:** Transactional claim with generation fencing and DB-clock lease authority.

**Implementation:** Select available work using SKIP LOCKED in short transactions; increment generation; set lease owner/deadline; insert attempt and event transactionally. Avoid database connection holding during guest execution.

**Acceptance:** Two claimers cannot both own one generation; stale generation rejected; claims on expired leases only after effect safety classification; expired owner cannot commit. Test with real concurrent DB connections.

**Evidence:** Race test including artificially delayed worker and SQL verification.
