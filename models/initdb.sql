CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  email STRING(100),
  password STRING(250),
  created_at TIMESTAMPZ NOT NULL current_timestamp()
);
