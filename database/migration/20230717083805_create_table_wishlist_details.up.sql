CREATE TABLE wishlist_details(
    product_id integer,
    customer_id UUID,
    Foreign Key (product_id) REFERENCES products(id),
    Foreign Key (customer_id) REFERENCES customers(id)
);