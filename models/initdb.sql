CREATE TABLE IF NOT EXISTS users (
  id         UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  email      STRING(100) UNIQUE NOT NULL,
  password   STRING(250) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp()
);
