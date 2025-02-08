-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create table calls(
    id serial primary key,
    session_id serial  references calls(id),
    environment_id serial references environments(id),
    name text not null,
    method text not null,
    path text,
    body text,
    query_params text,
    user_id text,
    called_at timestamp not null,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp not null
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table calls;
-- +goose StatementEnd
