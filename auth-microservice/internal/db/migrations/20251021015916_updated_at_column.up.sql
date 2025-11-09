CREATE OR REPLACE FUNCTION auth.updated_at_column()
RETURNS trigger AS $$
BEGIN
    NEW.updated_at := now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;