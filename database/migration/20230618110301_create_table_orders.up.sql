CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    id_product integer,
    id_customer UUID,
    quantity int,
    FOREIGN KEY (id_product) REFERENCES products(id),
    FOREIGN KEY (id_customer) REFERENCES customers(id)
);