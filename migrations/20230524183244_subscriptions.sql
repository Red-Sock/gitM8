-- +goose Up
-- +goose StatementBegin
CREATE TABLE subscriptions (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS identity,
    chat_id INTEGER,
    user_id INTEGER REFERENCES tg_users(tg_id),
    ticket_id INTEGER REFERENCES tickets(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subscriptions;
-- +goose StatementEnd
