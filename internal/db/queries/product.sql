-- name: GetProductByID :one
SELECT * FROM products WHERE id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListProductsByCategory :many
SELECT * FROM products
WHERE category = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CreateProduct :exec
INSERT INTO products (id, name, description, price, category, virtual_image_id, model_id)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateProduct :exec
UPDATE products
SET
    name             = $2,
    description      = $3,
    price            = $4,
    category         = $5,
    virtual_image_id = $6,
    model_id         = $7,
    updated_at       = now()
WHERE id = $1;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;
