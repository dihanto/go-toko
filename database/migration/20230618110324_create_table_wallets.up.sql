CREATE TABLE wallets(
    id serial PRIMARY KEY,
    id_customer UUID,
    balance integer,
    created_at integer,
    updated_at integer not NULL DEFAULT 0,
    deleted_at integer,
    FOREIGN KEY (id_customer) REFERENCES customers(id)
);
