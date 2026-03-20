-- +goose Up
CREATE TABLE IF NOT EXISTS products (
    id               TEXT PRIMARY KEY,
    name             TEXT        NOT NULL,
    description      TEXT        NOT NULL DEFAULT '',
    price            DOUBLE PRECISION NOT NULL DEFAULT 0,
    category         TEXT        NOT NULL DEFAULT '',
    virtual_image_id TEXT        NOT NULL DEFAULT '',
    model_id         TEXT        NOT NULL DEFAULT '',
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_products_category ON products (category);
CREATE INDEX IF NOT EXISTS idx_products_name     ON products USING gin (name gin_trgm_ops);

-- +goose Down
DROP TABLE IF EXISTS products;
