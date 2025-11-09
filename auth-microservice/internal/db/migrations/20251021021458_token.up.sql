CREATE TABLE IF NOT EXISTS auth."token" (
  id bigserial PRIMARY KEY,
  user_id bigint NOT NULL REFERENCES auth."user"(id),
  refresh_token varchar(512) NOT NULL UNIQUE,
  expires_at timestamptz NOT NULL,
  created_at timestamptz DEFAULT now(),
  revoked_at timestamptz
);