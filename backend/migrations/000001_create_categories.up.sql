CREATE TABLE categories (
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ,
    name        VARCHAR(100) NOT NULL,
    description TEXT,
    color       VARCHAR(7)   NOT NULL DEFAULT '#000000',
    CONSTRAINT  uq_categories_name UNIQUE (name)
);

CREATE INDEX idx_categories_deleted_at ON categories (deleted_at);
