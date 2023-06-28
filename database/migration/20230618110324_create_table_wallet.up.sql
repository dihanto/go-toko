CREATE TABLE wallet(
    id serial PRIMARY KEY,
    id_customer UUID,
    balance integer,
    FOREIGN KEY (id_customer) REFERENCES customers(id)
);
