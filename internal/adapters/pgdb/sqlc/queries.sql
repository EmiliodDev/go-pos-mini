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
