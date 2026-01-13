-- name: ListProducts :many
SELECT * FROM products
ORDER BY name;

-- name: FindProductByID :one
SELECT * FROM products
WHERE id = ?;

-- name: CreateOrder :execresult
INSERT INTO orders (
  customer_id
) VALUES (?);

-- name: CreateOrderItem :execresult
INSERT INTO order_items (
  order_id,
  product_id,
  quantity,
  price_cents
) VALUES (?, ?, ?, ?);
