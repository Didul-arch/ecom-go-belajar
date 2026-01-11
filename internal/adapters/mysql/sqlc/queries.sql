-- name: ListProducts :many
SELECT * FROM products
ORDER BY name;

-- name: FindProductByID :one
SELECT * FROM products
WHERE id = $1;

