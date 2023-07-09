CREATE TABLE banks (
    code SERIAL PRIMARY KEY,
    name VARCHAR(100),
    created_at integer,
    updated_at integer not NULL DEFAULT 0,
    deleted_at integer
);
