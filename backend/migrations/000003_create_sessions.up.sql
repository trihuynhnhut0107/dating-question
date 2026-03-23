CREATE TABLE sessions (
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    name       VARCHAR(200)
);

CREATE INDEX idx_sessions_deleted_at ON sessions (deleted_at);
