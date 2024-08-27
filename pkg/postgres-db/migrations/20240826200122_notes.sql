-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id          SERIAL PRIMARY KEY,
    user_name   TEXT NOT NULL,
    note        TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users
-- +goose StatementEnd