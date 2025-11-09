CREATE TABLE IF NOT EXISTS auth."user" (
  id bigserial PRIMARY KEY,
  login varchar(50) NOT NULL UNIQUE CHECK (char_length(login) >= 3),
  password_hash varchar(60) NOT NULL,
  is_active boolean DEFAULT true,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  deleted_at timestamptz
);

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_trigger WHERE tgname = 'set_user_updated_at') THEN
        CREATE TRIGGER set_user_updated_at
        BEFORE UPDATE ON auth."user"
        FOR EACH ROW
        EXECUTE FUNCTION auth.updated_at_column();
    END IF;
END$$;
