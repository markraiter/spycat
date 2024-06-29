CREATE TABLE IF NOT EXISTS cats (
    id                  SERIAL PRIMARY KEY,
    name                VARCHAR(255) NOT NULL UNIQUE,
    breed               VARCHAR(255) NOT NULL,
    years_of_experience INT,
    salary              INT,
    created_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_name ON cats (name);

CREATE OR REPLACE FUNCTION set_cats_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.updated_at IS NULL THEN
        NEW.updated_at = NOW();
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_cats_updated_at_trigger
BEFORE UPDATE ON cats
FOR EACH ROW
EXECUTE FUNCTION set_cats_updated_at();