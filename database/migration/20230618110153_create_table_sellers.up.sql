CREATE TABLE sellers (
    id UUID PRIMARY KEY,
    email VARCHAR(100) UNIQUE,
    name VARCHAR(100),
    password VARCHAR(100),
    registered_at integer,
    updated_at integer not NULL DEFAULT 0,
    deleted_at integer
);
