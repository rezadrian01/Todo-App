-- Table 2: todos
CREATE TABLE todos (
    id           SERIAL PRIMARY KEY,
    title        VARCHAR(255) NOT NULL,
    description  TEXT,
    completed    BOOLEAN      NOT NULL DEFAULT FALSE,
    priority     VARCHAR(10)  NOT NULL DEFAULT 'medium',
    due_date     TIMESTAMP,
    category_id  INT REFERENCES categories(id) ON DELETE SET NULL,
    created_at   TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP    NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX idx_todos_completed ON todos(completed);
CREATE INDEX idx_todos_category_id ON todos(category_id);
CREATE INDEX idx_todos_title_fts ON todos USING gin(to_tsvector('english', title));
