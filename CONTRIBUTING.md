# Contributing

Early architecture is being established. Start with an issue; changes to
execution guarantees, storage, or sandbox boundaries require an ADR.

Run `make check` before opening a pull request. Add tests for every new
state transition and any reported concurrency bug. Never submit guessed
performance results, real credentials, or unredacted user input.

Code review prioritizes correctness, small diffs, tests, explicit guarantees,
and measurable tradeoffs over framework count.
