-- +goose Up
CREATE TABLE IF NOT EXISTS orders (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  customer_id BIGINT NOT NULL,
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS order_items (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
  order_id BIGINT NOT NULL,
  product_id BIGINT NOT NULL,
  quantity INTEGER NOT NULL,
  price_cents INTEGER NOT NULL,
  CONSTRAINT fk_order FOREIGN KEY (order_id)
  REFERENCES orders(id)
);

-- +goose Down
DROP TABLE IF EXISTS order_items;

DROP TABLE IF EXISTS orders;

