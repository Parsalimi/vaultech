# CS Vault Technology Stack

**Status:** Proposed; requires approval because the committed baseline is Django/Python and the working tree is transitioning to Go.

| Layer | Proposed choice | Role | Alternatives considered |
|---|---|---|---|
| Frontend | React + TypeScript + Vite | Responsive browser application and typed API integration | Server-rendered templates, Next.js |
| Backend | Go modular monolith with REST/JSON | Authentication, authorization, domain services, persistence, and HTTP API | Existing Django prototype; a second backend is explicitly disallowed |
| Database | PostgreSQL | Relational constraints, transactions, indexes, and historical data | SQLite for local-only prototyping |
| Migrations | Versioned SQL migrations | Reproducible schema changes | ORM-generated migrations |
| Local development | Docker Compose for PostgreSQL and optional app services | Reproducible development environment | Host-installed services |
| Tests | Go unit/integration tests; frontend typecheck/build and component tests | Verify domain rules and critical workflows | Manual-only testing |

## Unresolved

- Approval of Go over the committed Django baseline.
- Go HTTP/router and database driver choices.
- Frontend styling and component-library choice.
- Supported Go, Node.js, PostgreSQL, and browser versions.
- Whether Docker Compose runs application containers or only PostgreSQL during development.

## Initial commands (after scaffolding)

Commands will be finalized with the implementation scaffold. The intended checks are `go test ./...`, frontend typecheck, frontend production build, and migration verification.
