-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    user_id    BIGSERIAL PRIMARY KEY,
    name       TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
