# Architecture

CS Vault will use a modular monolith: a React/TypeScript browser client calls a Go REST API, which contains domain services and PostgreSQL persistence. Backend modules are organized by capability (auth, academics, courses, professors, reviews) and HTTP handlers do not own business rules.

The first slice establishes configuration, a health endpoint, and authentication primitives. PostgreSQL-backed account activation and sessions are the next slice; the migration directory is versioned from the start.
