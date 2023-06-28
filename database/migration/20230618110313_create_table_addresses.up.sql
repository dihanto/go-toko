CREATE TABLE addresses(
    code VARCHAR(100) PRIMARY KEY,
    id_customer UUID,
    street VARCHAR(100),
    city VARCHAR(100),
    province VARCHAR(100),
    FOREIGN KEY (id_customer) REFERENCES customers(id)
);
