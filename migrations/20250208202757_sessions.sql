-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE sessions (
    id INTEGER PRIMARY KEY,
    environment_id INTEGER REFERENCES environments(id),
    user_id TEXT,
    tags TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE sessions;
-- +goose StatementEnd
