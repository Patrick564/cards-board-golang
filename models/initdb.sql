CREATE TABLE IF NOT EXISTS users (
  id         UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  email      STRING(100) NOT NULL UNIQUE,
  password   STRING(250) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp()
);

CREATE TABLE IF NOT EXISTS boards (
  id         UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  name       STRING(50),
  content    STRING(250),
  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp(),
  id_user    UUID NOT NULL REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS cards (
  id         UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  content    STRING(250),
  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp(),
  id_user    UUID NOT NULL REFERENCES users (id),
  id_board   UUID NOT NULL REFERENCES boards (id)
);
