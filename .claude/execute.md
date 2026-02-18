---
description: Implement feature per spec, keep diffs small, run Go tests.
---

# /execute (Go Microservices)

Feature: $ARGUMENTS

## Rules
- Must be on branch `feature/<slug>`.
- Follow clean architecture principles used in the project.
- Update `go.work` or `go.mod` if adding new modules/dependencies.
- Run tests before marking as complete.

## Steps
1) Confirm branch matches `feature/<slug>`.
2) Read requirements and identify impacted services.
3) Implement changes in small increments:
   - Domain/Entities
   - Repository/Use Cases
   - Handlers/API
4) Run service-specific tests:
   ```bash
   cd <service-name>
   go test ./...
   ```
5) Verify end-to-end flow if applicable (using Docker Compose or Skaffold).
6) Run linter/formatter (e.g., `go fmt ./...`, `go vet ./...`).
7) Set status to `In Progress` or `Ready for Review`.

## Output
- Summary of changes per service.
- Tests passed.
