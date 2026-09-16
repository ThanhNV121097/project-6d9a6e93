# Architecture overview

## Stack

| Part | Choice |
|---|---|
| Frontend | Next.js 15 App Router, TypeScript, Tailwind v3, ESLint |
| Backend | Go 1.22, `net/http`, pgx PostgreSQL driver |
| Data | PostgreSQL 16 |

## Layout

- `code/backend/cmd/api`: only executable; routes and boot sequence.
- `code/backend/migrations`: ordered embedded SQL migrations.
- `code/frontend/app`: App Router shell and frozen global tokens.
- `code/frontend/components`: one default-export component per story.
- `code/frontend/lib/mock`: UI-stage fixtures; backend stage replaces own fixture with API client.

## Contracts and conventions

- Backend reads `DATABASE_URL`, applies migrations, then listens on `PORT`, falling back to `APP_PORT`, then `8080`.
- `/healthz` returns 200 only after migration and database probe succeed.
- API routes use `/v1/...`; deployment proxy owns `/api` prefix.
- SQL uses pgx parameters. Migration versions are filename prefixes and are recorded in `schema_migrations`.
- One shared greeting row. Last completed update wins. No authentication.
- Frontend page is server composition root. Interactive story component begins with literal `"use client"` and default-exports its component.
- CSS modules use only tokens defined in `app/globals.css`; no token fallbacks or hardcoded visual values.

## Decisions

| Decision | Rejected | Tradeoff |
|---|---|---|
| pgx driver | ORM | Small direct SQL surface; no model abstraction for one row |
| self-applied SQL migrations | external migration job | Startup owns empty runtime DB setup; boot waits for DB |
| one singleton table row | greeting collection | Matches one shared value; expansion needs new schema/API |
| native `net/http` | router framework | Two routes need no dependency; add router when route count needs grouping |

## Environment

- Backend: `DATABASE_URL`, `PORT`, optional `APP_PORT` fallback.
- Frontend: `NEXT_PUBLIC_API_URL` for browser API calls; `API_ORIGIN` for server-side API calls.
- Compose: `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`; optional port and memory overrides in `docker-compose.yml`.

## Run

1. Copy `.env.example` to `.env` and change local values if needed.
2. Run `docker compose --profile local up --build` from repository root.
3. Open `http://localhost:3000`; backend health endpoint is `http://localhost:8080/healthz`.

Migrations run in filename order. `CREATE INDEX CONCURRENTLY` migrations run outside a transaction. Current migration seeds `Hello, World!`; changing seed after first boot does not overwrite saved data.
