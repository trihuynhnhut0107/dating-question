ALTER TABLE sessions ADD COLUMN IF NOT EXISTS uuid VARCHAR(36);
UPDATE sessions SET uuid = gen_random_uuid()::varchar WHERE uuid IS NULL;
ALTER TABLE sessions ALTER COLUMN uuid SET NOT NULL;
CREATE UNIQUE INDEX IF NOT EXISTS idx_sessions_uuid ON sessions(uuid);
