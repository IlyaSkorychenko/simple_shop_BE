CREATE TABLE IF NOT EXISTS products(
    id SERIAL PRIMARY KEY,
    name VARCHAR(256),
    price REAL CHECK (price > 0)
);
CREATE UNIQUE INDEX name_index ON products(name)