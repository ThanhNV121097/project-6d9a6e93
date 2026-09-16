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
