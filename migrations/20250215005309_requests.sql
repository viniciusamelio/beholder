-- +goose Up
-- +goose StatementBegin
create table requests (
    id serial primary key,
    environment_id SERIAL references environments(id),
    session_id SERIAL,
    name text not null,
    user_id text not null,
    method text not null,
    path text not null,
    body text,
    headers text,
    called_at datetime not null default current_timestamp,
    created_at datetime default current_timestamp,
    updated_at datetime
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table requests;
-- +goose StatementEnd
