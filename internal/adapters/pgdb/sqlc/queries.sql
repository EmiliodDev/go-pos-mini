-- name: ListProducts :many
SELECT
    product_id,
    name,
    description,
    price_cents,
    created_at,
    updated_at
FROM products;

-- name: GetProduct :one
SELECT
    product_id,
    name,
    description,
    price_cents,
    created_at,
    updated_at
FROM products
WHERE product_id = $1;

-- name: CreateProduct :one
INSERT INTO products (name, description, price_cents)
VALUES (
    $1, $2, $3
    ) RETURNING *;
