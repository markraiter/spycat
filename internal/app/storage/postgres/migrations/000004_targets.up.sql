CREATE TABLE IF NOT EXISTS targets (
    id         SERIAL PRIMARY KEY,
    mission_id INT NOT NULL REFERENCES missions(id) ON DELETE CASCADE,
    name       VARCHAR(255) NOT NULL,
    country    VARCHAR(255) NOT NULL,
    notes      TEXT,
    completed  BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE OR REPLACE FUNCTION set_targets_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.updated_at IS NULL THEN
        NEW.updated_at = NOW();
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_targets_updated_at_trigger
BEFORE UPDATE ON targets
FOR EACH ROW
EXECUTE FUNCTION set_targets_updated_at();