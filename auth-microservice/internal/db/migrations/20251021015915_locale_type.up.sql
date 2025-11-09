DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'locale_type' AND typnamespace = 'auth'::regnamespace) THEN
        CREATE TYPE auth.locale_type AS ENUM ('ru','en','kz');
    END IF;
END$$;