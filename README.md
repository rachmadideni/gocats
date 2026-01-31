# GoCats - Category & Product Service

GoCats is a small REST API for managing categories and products, implemented in Go.

**Key changes in this branch**

- Project refactored from a single `main.go` file into a layered architecture: `handlers`, `services`, `repository`, `models`/`domain`, `database`, and `config`.
- Using GORM as the ORM for database models and migrations.
- Using Viper for configuration loading (reads `.env` and environment variables) so values are not hard-coded.
- Migrations are safe to run on startup (non-fatal/idempotent): common "already exists" or prepared-statement warnings are logged but do not stop the server.

## Architecture

- `internal/config` — configuration loader using `viper` (reads `.env` or env vars).
- `internal/database` — database connection and migration utilities (GORM).
- `internal/models` / `internal/domain` — GORM models and response DTOs.
- `internal/repository` — DB access wrappers (preload relations where needed).
- `internal/services` — business logic and DTO mapping.
- `internal/handlers` — HTTP handlers exposing the API.

## Running Locally

1. Prerequisites: Go (1.24+), a PostgreSQL-compatible database (Supabase recommended).

2. Create a `.env` file in the project root (you can copy `.env.example`) and set `DATABASE_URL`, `SERVER_HOST`, `SERVER_PORT`.

3. Run the application:

```bash
go run main.go
```

Notes:
- The app uses `viper` to load `.env` automatically. If `DATABASE_URL` is missing the app will return a helpful error.
- The DB connection will be opened with settings that avoid prepared-statement caching issues (useful for pooled Supabase connections).

## API Endpoints

- `GET /health` — Health check
- `GET /api/categories` — List categories
- `GET /api/categories/{id}` — Get category by ID
- `POST /api/categories` — Create category
- `PUT /api/categories/{id}` — Update category
- `DELETE /api/categories/{id}` — Delete category

- `GET /api/products` — List products (response includes nested `category` object)
- `GET /api/products/{id}` — Get product by ID (includes nested `category`)
- `GET /api/products?category_id={id}` — Filter products by category
- `POST /api/products` — Create product
- `PUT /api/products/{id}` — Update product
- `DELETE /api/products/{id}` — Delete product

## Migrations

- Migrations are executed on startup by default. They are idempotent and non-fatal; the server will continue to run if a table already exists or if there are benign prepared-statement warnings.

## Example Requests

Use the provided `request.http` (for VS Code REST Client) or curl. Example:

```bash
# Health check
curl http://localhost:6000/health

# Get all products (category nested)
curl http://localhost:6000/api/products

# Create a product
curl -X POST http://localhost:6000/api/products \
  -H "Content-Type: application/json" \
  -d '{"name":"iPhone 15 Pro","description":"Phone","price":15999000,"stock":50,"category_id":1}'
```

## Notes & Troubleshooting

- If you see prepared-statement cache errors when talking to a pooled Postgres (e.g., Supabase), the app disables prepared-statement caching and uses the simple protocol where needed.
- If you see prepared-statement cache errors when talking to a pooled Postgres (e.g., Supabase), the app disables prepared-statement caching and uses the simple protocol where needed.
  - Related discussion: https://forum.bubble.io/t/sql-connector-issue-prepared-statement-supabase-integration/303849/4
- Ensure `DATABASE_URL` is correct and the DB is reachable before starting the app.

---


