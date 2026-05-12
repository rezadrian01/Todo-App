DROP INDEX IF EXISTS idx_todos_title_fts;
CREATE INDEX idx_todos_title_desc_fts ON todos USING gin(to_tsvector('english', title || ' ' || COALESCE(description, '')));
