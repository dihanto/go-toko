CREATE TABLE addresses(
    code VARCHAR(100) PRIMARY KEY,
    id_customer UUID,
    street VARCHAR(100),
    city VARCHAR(100),
    province VARCHAR(100),
    created_at integer,
    updated_at integer not NULL DEFAULT 0,
    deleted_at integer,
    FOREIGN KEY (id_customer) REFERENCES customers(id)
);
