CREATE TABLE IF NOT EXISTS users (
  id         UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  username   STRING(100) NOT NULL UNIQUE,
  email      STRING(100) NOT NULL UNIQUE,
  password   STRING(250) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp()
);

CREATE TABLE IF NOT EXISTS boards (
  id         UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  name       STRING(50) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp(),
  user_id    UUID NOT NULL REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS cards (
  id         UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  creator    STRING(100) NOT NULL DEFAULT 'anonymous',
  content    STRING(250) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp(),
  user_id    UUID NOT NULL REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
  board_id   UUID NOT NULL REFERENCES boards (id) ON UPDATE CASCADE ON DELETE CASCADE
);
