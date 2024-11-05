-- Active: 1730083286169@@127.0.0.1@5432@serviceinventory@public
-- Buat tabel kategori barang
CREATE TABLE Users (
    user_id VARCHAR(10) PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE categories (
    id VARCHAR(10) PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

SELECT * FROM categories LIMIT 3 OFFSET 1;


-- Buat tabel lokasi penyimpanan barang
CREATE TABLE locations (
    id VARCHAR(10) PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Buat tabel produk/barang
CREATE TABLE products (
    id VARCHAR(10) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    category_id VARCHAR(10) REFERENCES categories(id),
    location_id VARCHAR(10) REFERENCES locations(id)
);

SELECT 
    p.id AS product_id, 
    p.name AS product_name, 
    p.price AS product_price, 
    p.stock AS product_stock, 
    c.name AS category_name, 
    l.name AS location_name
FROM 
    products p
JOIN 
    categories c ON c.id = p.category_id
JOIN 
    locations l ON l.id = p.location_id
WHERE 
    -- c.name LIKE '%Elektronik%'
    -- p.name LIKE '%Scanner%'
    p.id LIKE '%%pd%%'
LIMIT 2 OFFSET 0;

-- Buat tabel transaksi
CREATE TABLE transactions (
    id VARCHAR(10) PRIMARY KEY,
    product_id VARCHAR(10) REFERENCES products(id),
    quantity INT NOT NULL,
    transaction_type VARCHAR(10) CHECK (transaction_type IN ('in', 'out')),
    information VARCHAR(256),
    timestamp TIMESTAMPTZ DEFAULT NOW()
);

SELECT trx.id, trx.quantity, trx.transaction_type AS "Trx Tipe", trx.information, 
       p.name AS Items, c.name AS Category, l.name AS Location
FROM transactions trx
JOIN products p ON trx.product_id = p.id
JOIN categories c ON p.category_id = c.id
JOIN locations l ON p.location_id = l.id
LIMIT 5 OFFSET 2;


-- Masukkan data dummy
-- Masukan dara user
INSERT INTO Users (user_id, username, password) VALUES
('adm1', 'admin1', 'admin123');

SELECT * FROM users;

-- Masukkan data kategori barang
INSERT INTO categories (id, name) VALUES 
    ('ctg1', 'Elektronik'), 
    ('ctg2', 'Bahan Baku'),
    ('ctg3', 'Peralatan Kantor'),
    ('ctg4', 'Alat Tulis'),
    ('ctg5', 'Furnitur');

-- Masukkan data lokasi penyimpanan barang
INSERT INTO locations (id, name) VALUES 
    ('loc1', 'Gudang Utama Rak Nomer 1'), 
    ('loc2', 'Gudang Cabang Rak Nomer 2'),
    ('loc3', 'Gudang Lantai 1 Rak Nomer 3'),
    ('loc4', 'Gudang Lantai 2 Rak Nomer 4');

TRUNCATE locations;

SELECT * FROM locations;

-- Masukkan data produk/barang
INSERT INTO products (id, name, price, stock, category_id, location_id) VALUES 
    ('pd1', 'Laptop', 5000000.00, 20, 'ctg1', 'loc1'),
    ('pd2', 'Kabel Listrik', 5000.00, 50, 'ctg2', 'loc2'),
    ('pd3', 'Printer', 1500000.00, 10, 'ctg1', 'loc1'),
    ('pd4', 'Monitor', 2000000.00, 15, 'ctg1', 'loc3'),
    ('pd5', 'Meja Kantor', 300000.00, 8, 'ctg5', 'loc4'),
    ('pd6', 'Kursi Kantor', 150000.00, 25, 'ctg5', 'loc4'),
    ('pd7', 'Pensil', 500.00, 100, 'ctg4', 'loc2'),
    ('pd8', 'Pulpen', 1000.00, 200, 'ctg4', 'loc2'),
    ('pd9', 'Notebook', 10000.00, 150, 'ctg4', 'loc1'),
    ('pd10', 'Kertas A4', 50000.00, 300, 'ctg4', 'loc1'),
    ('pd11', 'Scanner', 2500000.00, 5, 'ctg1', 'loc3'),
    ('pd12', 'Proyektor', 7000000.00, 3, 'ctg1', 'loc1'),
    ('pd13', 'Mouse', 15000.00, 60, 'ctg1', 'loc3'),
    ('pd14', 'Keyboard', 25000.00, 45, 'ctg1', 'loc3'),
    ('pd15', 'Saklar', 2000.00, 100, 'ctg2', 'loc2'),
    ('pd16', 'Lampu LED', 10000.00, 80, 'ctg2', 'loc2'),
    ('pd17', 'Rak Besi', 500000.00, 6, 'ctg5', 'loc4'),
    ('pd18', 'Komputer', 4000000.00, 7, 'ctg1', 'loc1'),
    ('pd19', 'Switch Hub', 350000.00, 12, 'ctg1', 'loc3'),
    ('pd20', 'Kabel HDMI', 15000.00, 90, 'ctg2', 'loc1');

-- Masukkan data transaksi
INSERT INTO transactions (id, product_id, quantity, transaction_type, information, timestamp) VALUES
    ('trx1', 'pd1', 5, 'in', 'Restocked new items', '2024-01-01 10:00:00'),
    ('trx2', 'pd1', 2, 'out', 'Customer purchase', '2024-01-02 14:30:00'),
    ('trx3', 'pd2', 10, 'in', 'Restocked items', '2024-01-03 09:15:00'),
    ('trx4', 'pd2', 5, 'out', 'Order for project', '2024-01-04 16:45:00'),
    ('trx5', 'pd3', 4, 'in', 'Supplier delivery', '2024-01-05 11:00:00'),
    ('trx6', 'pd3', 3, 'out', 'Customer return', '2024-01-06 13:20:00'),
    ('trx7', 'pd4', 2, 'in', 'Inventory refill', '2024-01-07 10:10:00'),
    ('trx8', 'pd4', 1, 'out', 'Product sold', '2024-01-08 15:50:00'),
    ('trx9', 'pd5', 3, 'in', 'New shipment received', '2024-01-09 08:40:00'),
    ('trx10', 'pd5', 2, 'out', 'Sample for client', '2024-01-10 17:30:00'),
    ('trx11', 'pd6', 10, 'in', 'Regular stock update', '2024-01-11 12:00:00'),
    ('trx12', 'pd6', 5, 'out', 'Sold in bulk order', '2024-01-12 16:00:00'),
    ('trx13', 'pd7', 20, 'in', 'Year-end restock', '2024-01-13 09:45:00'),
    ('trx14', 'pd8', 30, 'in', 'Emergency stock', '2024-01-14 10:30:00'),
    ('trx15', 'pd9', 15, 'out', 'Clearance sale', '2024-01-15 14:20:00'),
    ('trx16', 'pd10', 50, 'in', 'Bulk purchase from supplier', '2024-01-16 11:30:00'),
    ('trx17', 'pd11', 1, 'in', 'Special order for display', '2024-01-17 13:15:00'),
    ('trx18', 'pd12', 1, 'in', 'One-time stock', '2024-01-18 10:05:00'),
    ('trx19', 'pd13', 10, 'in', 'Holiday sale restock', '2024-01-19 09:25:00'),
    ('trx20', 'pd14', 5, 'out', 'Donated to charity', '2024-01-20 16:10:00');


SELECT * FROM transactions;