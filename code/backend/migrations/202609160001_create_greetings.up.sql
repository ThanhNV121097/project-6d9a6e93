CREATE TABLE IF NOT EXISTS schema_migrations (
  version text PRIMARY KEY,
  applied_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE greetings (
  id boolean PRIMARY KEY CHECK (id),
  text text NOT NULL CHECK (btrim(text) <> ''),
  updated_at timestamptz NOT NULL
);

INSERT INTO greetings (id, text, updated_at) VALUES (true, 'Hello, World!', now())
ON CONFLICT (id) DO NOTHING;
