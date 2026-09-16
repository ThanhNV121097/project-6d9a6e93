# ERD

## `greetings`

One shared persisted greeting.

| Column | PostgreSQL type | Rules |
|---|---|---|
| `id` | `boolean` | Primary key, constrained to `true`; singleton row |
| `text` | `text` | Required; trimmed non-empty value |
| `updated_at` | `timestamptz` | Required; UTC update time |

Seed row: `id = true`, `text = 'Hello, World!'`.

## `schema_migrations`

Migration runner bookkeeping.

| Column | PostgreSQL type | Rules |
|---|---|---|
| `version` | `text` | Primary key; migration filename prefix |
| `applied_at` | `timestamptz` | Required; application time |

## Relationships

No foreign keys. `greetings` has exactly one row. `schema_migrations` is operational metadata and has no product relationship.

## Migration plan — Persisted editable greeting

Forward migration creates `greetings` with `id boolean PRIMARY KEY CHECK (id)`, `text text NOT NULL CHECK (btrim(text) <> '')`, and `updated_at timestamptz NOT NULL`. It inserts singleton seed row with `INSERT ... ON CONFLICT (id) DO NOTHING`, preserving existing visitor changes on restart or re-run. It also creates `schema_migrations` if absent.

Backward migration drops `greetings`; this destroys saved greeting data and is only safe before production data matters. Forward migration is safe on populated databases: it neither changes existing tables nor overwrites singleton row. No extra index is needed: reads and writes address primary-key `id = true`.
