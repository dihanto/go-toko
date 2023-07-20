CREATE TABLE order_details(
    id  SERIAL PRIMARY KEY,
    id_product INT,
    quantity INT,
    id_order INT,
    Foreign Key (id_order) REFERENCES orders(id),
    Foreign Key (id_product) REFERENCES products(id)
)