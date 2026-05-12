-- Table 1: categories
CREATE TABLE categories (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100) NOT NULL UNIQUE,
    color       VARCHAR(7)   NOT NULL DEFAULT '#3B82F6',
    created_at  TIMESTAMP    NOT NULL DEFAULT NOW()
);
