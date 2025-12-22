CREATE TABLE IF NOT EXISTS auth."session" (
  id bigserial PRIMARY KEY,
  user_id bigint NOT NULL REFERENCES auth."user"(id),
  ip_address inet NOT NULL,
  user_agent varchar(1024),
  login_at timestamptz DEFAULT now(),
  logout_at timestamptz,
  last_seen_at timestamptz,
  is_active boolean DEFAULT true
);