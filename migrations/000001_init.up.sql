CREATE TABLE IF NOT EXISTS storages (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    is_available BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    size INTEGER CHECK (size >= 0) NOT NULL
);

CREATE TABLE IF NOT EXISTS products_to_storages (
    storage_id INTEGER REFERENCES storages (id),
    product_id INTEGER REFERENCES products (id),
    amount INTEGER CHECK (amount >= 0) NOT NULL,
    PRIMARY KEY (storage_id, product_id)
);

CREATE TABLE IF NOT EXISTS reservations (
    storage_id INTEGER REFERENCES storages (id),
    product_id INTEGER REFERENCES products (id),
    client_id INTEGER NOT NULL,
    amount INTEGER CHECK (amount >= 0) NOT NULL,
    PRIMARY KEY (storage_id, product_id, client_id)
);

-- According to the requirements, the service should not support adding stores and products --
-- So we have to add them with migration --

INSERT INTO storages (id, name, is_available)
VALUES (1, 'MSK', true), (2, 'SPB', true), (3, 'Kirov', false);

INSERT INTO products (id, name, size)
VALUES (1, 'Shoes', 10), (2, 'Oven', 100), (3, 'towel', 20);

INSERT INTO products_to_storages (storage_id, product_id, amount)
VALUES (1, 1, 100), (1, 3, 10), (2, 1, 50), (2, 2, 30), (3, 1, 200), (3, 3, 1000);