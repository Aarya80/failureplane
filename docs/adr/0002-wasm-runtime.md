# ADR-0002 — Wasmtime Component Model as the first guest runtime

Status: **Proposed; requires security and compatibility spike**
Date: 2026-10-11

## Decision
Use a Rust host with Wasmtime and a WIT-defined capability boundary. Guest
components must declare imports and are denied unrestricted host networking,
filesystem, environment, clocks and randomness by default. A separate broker
handles approved external effects.

## Why not Docker/gVisor/Firecracker first?
Containers and microVMs support more existing executables but add host/container
lifecycle, image distribution, cold-start, kernel and runtime-isolation
complexity. WASM components give a narrower, inspectable application ABI.

## Limits of the decision
WASM is not a general replacement for Linux process execution. Language,
filesystem, networking and threading compatibility may restrict adopters.
Running WASM safely still requires patched runtimes, host process privilege
reduction, strict resource limits, input validation, and adversarial tests.

## Spike acceptance
- Run one deterministic component with input and bounded output.
- Reject component with prohibited imports.
- Bound CPU/fuel, memory, execution wall time and output size.
- Report the Wasmtime version and security posture.

Do not implement this in the first Go state-machine session.
