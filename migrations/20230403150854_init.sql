-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tg_users
(
    tg_id    INTEGER PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS tickets
(
    id         INTEGER PRIMARY KEY GENERATED ALWAYS AS identity ( increment by 1 start 1),
    name       TEXT,
    owner_id   INTEGER REFERENCES tg_users (tg_id),
    uri        TEXT,
    git_system INTEGER,
    is_private BOOL
);

CREATE TABLE IF NOT EXISTS ticket_rules
(
    id        INTEGER PRIMARY KEY GENERATED ALWAYS AS identity ( increment by 1 start 1),
    ticket_id INTEGER REFERENCES tickets (id),
    payload   bytea,
    rule_type INTEGER
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ticket_rules;
DROP TABLE IF EXISTS tickets;
DROP TABLE IF EXISTS tg_users;
-- +goose StatementEnd
