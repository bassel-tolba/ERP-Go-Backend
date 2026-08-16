<img width="800" height="500" alt="Untitled design (1)" src="https://github.com/user-attachments/assets/c27ebf45-69a3-44e9-85ad-40ef1397970a" />
# Go-ERP

Go-based ERP system. Currently laying down the foundational plumbing—auth, role-based access control, and multi-tenant company structures—before diving into the actual business logic (accounting, inventory, etc.). Building it with strict typing and schema-first design from day one so it doesn't collapse on itself later. It happens.

## Stack

- **Backend:** Go (1.22+), standard `net/http` mux.
- **API Setup:** OpenAPI 3 schema-first, generated strictly via `oapi-codegen`.
- **DB:** PostgreSQL (`pgxpool`), `sqlc` for type-safe queries, `goose` for migrations.
- **Auth:** JWT (v5) + bcrypt.
- **Other:** `viper` for config management, `slog` for structured logging, custom OpenAPI validation middlewares.

## What's actually in here

This isn't a toy CRUD app, but it *is* still the foundation. The `modules` currently cover:

**Auth & Security**
- JWT issuance, password hashing, expiration handling.
- Strict OpenAPI request/response validation middleware (malformed payloads get rejected before hitting the handlers).

**User Management & RBAC**
- Users, Roles, and Permissions.
- Many-to-many relationship mapping (Assign roles to users, assign permissions to roles).

**Core Setup (Multi-tenant prep)**
- `Companies` module — basic CRUD for company profiles.

**Tooling**
- `traverse.go`: A custom context-gathering script to dump the codebase for AI analysis.
- `sqlc.yaml` & `openapi-config.yaml` to regenerate the tedious parts of the codebase.

## Getting started

You'll need a local Postgres instance running. 

1. Update `config.env` with your actual database credentials (don't just use my `bassel:flstudio` local setup).
2. Run the database migrations:
```bash
goose -dir internal/modules/base/migrations postgres "postgres://bassel:flstudio@localhost:5432/go-erp2?sslmode=disable" up
```
3. Fire up the server:
```bash
go run main.go
```

If you change the SQL queries or the OpenAPI spec, you'll need to regenerate the code:
```bash
sqlc generate
go generate ./...
```

## Tests

Since this is an API-first backend, there's a custom black-box testing script `test.py` in the root. It hammers the API endpoints to validate both happy paths (200/201/204) and unhappy paths (400/404/422 constraint violations). Run it with:

```bash
python3 test.py
```

## Known rough edges

Being upfront:
- **No frontend yet.** It's just a raw API right now. Bring your own UI.
