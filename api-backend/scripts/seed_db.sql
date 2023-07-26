-- Inventory table
CREATE TABLE IF NOT EXISTS inventory (
    id SERIAL PRIMARY KEY,
    product_name TEXT NOT NULL,
    quantity INTEGER, 
    price NUMERIC(10,2)
);

-- Delete to clear seed info
DELETE FROM inventory;

-- Seed data
INSERT INTO inventory (product_name, quantity, price) VALUES ('Notebook', 20, 9.99);
INSERT INTO inventory (product_name, quantity, price) VALUES ('Pen', 100, 2.99);
INSERT INTO inventory (product_name, quantity, price) VALUES ('Mechanical Pencil', 50, 6.99);
INSERT INTO inventory (product_name, quantity, price) VALUES ('Bento Box', 10, 19.99);
INSERT INTO inventory (product_name, quantity, price) VALUES ('Tote Bag', 15, 29.99);