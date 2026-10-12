# Threat model and security work (design, not verified)

## Assets
Execution/tenant data, API credentials, provider secrets, effect logs, host
resources, PostgreSQL data, artifacts, and replay evidence.

## Adversaries
Malformed API client; malicious WebAssembly guest; one tenant attempting to
read another's records; untrusted external service responses; compromised
worker; hostile traffic during public demo.

## Required control families
- Scoped hashed client keys initially; worker authentication for remote nodes.
- Tenant-qualified SQL; PostgreSQL RLS once auth and tenant context exist.
- Immutable content-addressed component artifacts.
- No guest default networking/filesystem/clock/randomness/env imports.
- Broker resolves capability IDs server-side; allowlists targets and protects
  internal IPs, DNS rebinding and metadata endpoints; no secrets to guest.
- Fuel, memory, wall-time, input and output caps; dedicated low-privilege hosts.
- Secrets in environment/secret store, never in replay artifacts.
- Authenticated read controls and retention/redaction for execution logs.
- Deny arbitrary guest uploads on public demo until separate review.

## Security tests
Malicious oversized input, forbidden import, infinite loop, memory abuse,
unauthorized tenant reads, cross-tenant execution IDs, SSRF, malformed artifact,
secret leakage in traces, unauthorized replay, artifact swap/tampering.

## Explicit caveat
The provided starter does NOT execute untrusted code. Do not describe it as
secure sandbox infrastructure until the Rust host is implemented and audited.
