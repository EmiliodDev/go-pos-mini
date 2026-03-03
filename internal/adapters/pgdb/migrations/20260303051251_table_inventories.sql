-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS inventories (
    inventory_id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL,
    warehouse_id BIGINT NOT NULL,
    stock BIGINT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS inventories;
-- +goose StatementEnd
