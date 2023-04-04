-- +goose Up
-- +goose StatementBegin


CREATE TABLE IF NOT EXISTS tg_users
(
    id    INTEGER PRIMARY KEY GENERATED ALWAYS AS identity ( increment by 1 start 1),
    tg_id INTEGER UNIQUE
);

CREATE TABLE IF NOT EXISTS tickets
(
    id       INTEGER PRIMARY KEY GENERATED ALWAYS AS identity ( increment by 1 start 1),
    name     TEXT,
    owner_id INTEGER REFERENCES tg_users (id),
    web_url      TEXT
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tickets;
DROP TABLE IF EXISTS tg_users;
-- +goose StatementEnd
