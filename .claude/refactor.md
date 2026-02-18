---
description: Refactor safely with tests as safety net. Updates technical-debt.md and feature spec if refactor is part of feature scope.
---

# /refactor (Feature-First)

Feature slug + refactor goal: $ARGUMENTS

## Rules
- No behavior changes unless the feature spec explicitly allows it.
- Ensure tests exist; add targeted tests if missing.
- Update docs:
  - `docs/ai/technical-debt.md` (resolved items)
  - `docs/features/<slug>.md` if refactor is within scope

## Steps
1) Read feature spec + identify refactor boundary.
2) Ensure safety net tests (`go test`).
3) Refactor in small commits.
4) Run lint/tests (`go vet`, `go test`).
5) Update docs and summarize.
