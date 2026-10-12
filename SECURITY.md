# Security policy

FailurePlane is an early engineering experiment, **not production ready**.
Do not run arbitrary untrusted user components or expose it publicly.

For suspected vulnerabilities, do not publish exploits or credentials in a
public issue. Open a private GitHub security advisory on the repository after
its owner enables this feature, or contact the maintainer through an official
verified project channel when available.

The planned Wasmtime runtime will require explicit threat modeling, pinned
patched dependencies and isolation tests before any untrusted-code claims.
