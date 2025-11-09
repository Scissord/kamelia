CREATE TABLE IF NOT EXISTS auth.profile (
  id bigserial PRIMARY KEY,
  user_id bigint NOT NULL,
  first_name_encrypted BYTEA,
  first_name_hash varchar(64),
  last_name_encrypted BYTEA,
  last_name_hash varchar(64),
  middle_name_encrypted BYTEA,
  middle_name_hash varchar(64),
  email_encrypted BYTEA,
  email_hash varchar(64) UNIQUE,
  phone_encrypted BYTEA,
  phone_hash varchar(64) UNIQUE,
  avatar_url varchar(512),
  birthday_encrypted BYTEA,
  birthday_hash varchar(64),
  gender auth.gender_type DEFAULT 'other'::auth.gender_type,
  locale auth.locale_type DEFAULT 'ru'::auth.locale_type,
  timezone varchar(50),
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  CONSTRAINT profile_user_id_key UNIQUE (user_id),
  CONSTRAINT profile_user_id_fkey FOREIGN KEY (user_id) REFERENCES auth."user"(id) ON DELETE CASCADE
);

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_trigger WHERE tgname = 'set_profile_updated_at') THEN
        CREATE TRIGGER set_profile_updated_at
        BEFORE UPDATE ON auth.profile
        FOR EACH ROW
        EXECUTE FUNCTION auth.updated_at_column();
    END IF;
END$$;