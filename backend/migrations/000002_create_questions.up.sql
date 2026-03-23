CREATE TYPE difficulty AS ENUM ('easy', 'medium', 'hard');

CREATE TABLE questions (
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ,
    text        TEXT        NOT NULL,
    category_id INTEGER     NOT NULL REFERENCES categories (id),
    difficulty  difficulty  NOT NULL DEFAULT 'easy',
    is_active   BOOLEAN     NOT NULL DEFAULT TRUE
);

CREATE INDEX idx_questions_deleted_at  ON questions (deleted_at);
CREATE INDEX idx_questions_category_id ON questions (category_id);
