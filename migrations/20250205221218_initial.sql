-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';
CREATE TABLE environments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    tags TEXT[],
    base_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    environment_id INTEGER REFERENCES environments(id),
    user_id TEXT,
    description TEXT,
    tags TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE calls(
    id SERIAL PRIMARY KEY,
    session_id INTEGER REFERENCES sessions(id),
    name TEXT NOT NULL,
    method TEXT NOT NULL,
    path TEXT,
    headers TEXT,
    body TEXT,
    query_params TEXT,
    called_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
DROP TABLE calls;
DROP TABLE sessions;
DROP TABLE environments;
-- +goose StatementEnd
