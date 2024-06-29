CREATE TABLE IF NOT EXISTS missions (
    id                   SERIAL PRIMARY KEY,
    cat_id               INT REFERENCES cats(id) ON DELETE CASCADE,
    notes                TEXT,
    completed            BOOLEAN NOT NULL DEFAULT FALSE,
    created_at           TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at           TIMESTAMP
);

CREATE OR REPLACE FUNCTION set_missions_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.updated_at IS NULL THEN
        NEW.updated_at = NOW();
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_missions_updated_at_trigger
BEFORE UPDATE ON cats
FOR EACH ROW
EXECUTE FUNCTION set_missions_updated_at();