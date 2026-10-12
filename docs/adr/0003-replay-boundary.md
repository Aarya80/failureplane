# ADR-0003 — Logical replay and unknown external effects

Status: **Accepted design boundary; implementation pending**
Date: 2026-10-11

## Decision
Replay means **logical re-execution of a pinned tool artifact and recorded
host-boundary responses**. It does not reproduce native thread interleavings,
physical time, or unrecorded external state. The replay engine must deny live
external writes and report INCOMPLETE when required evidence is missing.

Persist an effect-intent record before dispatching an external request. A crash
between downstream acceptance and local response commit can leave UNKNOWN.

- With a downstream idempotency key, retry/reconcile according to provider docs.
- With downstream status lookup, reconcile before further effects.
- Without either, stop automatic retries and surface reconciliation required.

A fenced result commit protects local state, not arbitrary downstream effects.

## Required experiment
Simulate a provider accepting a write then dropping the response. Kill the
worker. Verify that no second write is blindly dispatched without an explicit
safe policy. Inspect effects and event history after restart.
