CREATE TABLE session_questions (
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ,
    session_id  INTEGER     NOT NULL REFERENCES sessions  (id),
    question_id INTEGER     NOT NULL REFERENCES questions (id),
    is_answered BOOLEAN     NOT NULL DEFAULT FALSE,
    answered_at TIMESTAMPTZ
);

CREATE INDEX idx_session_questions_deleted_at  ON session_questions (deleted_at);
CREATE INDEX idx_session_questions_session_id  ON session_questions (session_id);
CREATE INDEX idx_session_questions_question_id ON session_questions (question_id);
