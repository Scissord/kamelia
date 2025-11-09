DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_type' AND typnamespace = 'auth'::regnamespace) THEN
        CREATE TYPE auth.gender_type AS ENUM ('male','female','other');
    END IF;
END$$;