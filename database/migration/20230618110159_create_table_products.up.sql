CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    id_seller UUID,
    name VARCHAR(100),
    price integer,
    quantity integer,
    created_at integer,
    updated_at integer not NULL DEFAULT 0,
    deleted_at integer,
    FOREIGN KEY (id_seller) REFERENCES sellers(id)
);
