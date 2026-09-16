# Service contracts

Backend receives paths after deployment proxy strips `/api`. All paths below intentionally omit that prefix.

## Shared error envelope

```json
{"error":{"code":"VALIDATION_FAILED","message":"Greeting must not be empty."}}
```

| HTTP | Code | Message | When |
|---|---|---|---|
| 400 | `MALFORMED_REQUEST` | `Request body is malformed.` | Bad JSON, wrong field type, or unknown field |
| 422 | `VALIDATION_FAILED` | `Greeting must not be empty.` | `text` trims to empty |
| 500 | `INTERNAL` | `Internal server error.` | Query or unexpected server failure |
| 503 | `UNAVAILABLE` | `Service unavailable.` | Database dependency refused or unavailable |

## Endpoints

### `GET /v1/greeting`

Returns current shared greeting.

Auth: none; public Visitor access.

Response `200`:

```json
{"text":"Hello, World!"}
```

Errors: `500 INTERNAL`, `503 UNAVAILABLE`.

### `PUT /v1/greeting`

Replaces shared greeting. Server trims leading and trailing whitespace before storing. Last completed save wins.

Auth: none; public Visitor access.

Request:

```json
{"text":"Pipeline accepted"}
```

Response `200`:

```json
{"text":"Pipeline accepted"}
```

Errors: `400 MALFORMED_REQUEST`, `422 VALIDATION_FAILED`, `500 INTERNAL`, `503 UNAVAILABLE`.

### `GET /healthz`

Operational endpoint. Returns `200` with `{"status":"ok"}` only after migrations and `SELECT 1` succeed; otherwise `503` with shared `UNAVAILABLE` envelope.

## Story design note — Persisted editable greeting

UI mock exports `Greeting` as `{ "text": string }`. Contract matches it exactly for both read and save responses, so backend swap only replaces mock import/client behavior. Mock has no error shape; shared envelope above governs API errors. Frontend keeps client-side empty validation for focus and `Enter a greeting.` message, while backend independently rejects bypassed invalid requests.
