-- Active: 1730083286169@@127.0.0.1@5432
CREATE DATABASE latihan13;


CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50),
    age INT,
    email VARCHAR(100),
    birth_date DATE,
    registration_date TIMESTAMP
 );
 
 
DROP TABLE users;


INSERT INTO users (username, age, email, birth_date, registration_date) VALUES
('jhon_doe', 30, 'jhon.doe@example.com', '1990-01-15', NOW()),
('jhon_doe1', NULL, 'jhon.doe@example.com', '1990-01-15', NOW()),
(NULL, 32, 'jhon.doe@example.com', '1990-01-15', NOW()),
('jhon_doe3', 33, NULL, '1990-01-15', NOW()),
('jhon_doe4', 34, 'jhon.doe@example.com', NULL, NOW()),
('jhon_doe5', 35, 'jhon.doe@example.com', '1990-01-15', NULL);



CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    email VARCHAR(100)
 );
 
INSERT INTO customers (username, password, email) VALUES
('admin', 'adminpassword', 'admin@excample.com'),
('customer1', 'password1', 'customer1@excample.com'),
('customer2', 'password2', 'customer2@excample.com'),
('customer3', 'password3', 'customer3@excample.com');


SELECT * FROM customers;


CREATE TABLE products ( 
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    price DECIMAL(10, 2) NOT NULL
);


CREATE TABLE orders ( 
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE order_items ( 
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);


INSERT INTO products (name, stock, price) VALUES
('P00D', 10, 55.00),
('P00E', 20, 30.00),
('P00F', 15, 20.00);







CREATE TABLE Customers (
    customer_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Drivers (Data Master)
CREATE TABLE Drivers (
    driver_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    vehicle_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Locations (Data Master)
CREATE TABLE Locations (
    location_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    type VARCHAR(50) CHECK (type IN ('pickup', 'dropoff'))
);

-- Tabel Orders (Data Transaksi)
CREATE TABLE Orders (
    order_id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL REFERENCES Customers(customer_id) ON DELETE CASCADE,
    driver_id INT REFERENCES Drivers(driver_id) ON DELETE SET NULL,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    pickup_location INT NOT NULL REFERENCES Locations(location_id) ON DELETE RESTRICT,
    dropoff_location INT NOT NULL REFERENCES Locations(location_id) ON DELETE RESTRICT,
    total_fare DECIMAL(10, 2) NOT NULL
);

-- Tabel OrderStatus (Data Transaksi)
CREATE TABLE OrderStatus (
    status_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES Orders(order_id) ON DELETE CASCADE,
    status VARCHAR(50) CHECK (status IN ('ongoing', 'completed', 'cancelled')),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel CustomerLogins (Data Transaksi)
CREATE TABLE CustomerLogins (
    login_id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL REFERENCES Customers(customer_id) ON DELETE CASCADE,
    login_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    logout_time TIMESTAMP
);

-- Insert query

-- add customers
INSERT INTO Customers (name, email, phone_number) VALUES
('Haidar Silalahi', 'haidar@gmail.com', '081234567890'),
('Darmi Agustina', 'darmi@gmail.com', '081234567891'),
('Darto Kurniawan', 'darto@gmail.com', '081234567892');


-- add logins
INSERT INTO CustomerLogins (customer_id, login_time, logout_time) VALUES
(1, '2024-10-01 07:00:00', '2024-10-01 10:00:00'),
(2, '2024-11-05 17:30:00', NULL),
(3, '2024-10-20 08:30:00', '2024-10-20 11:00:00');


-- add drivers
INSERT INTO Drivers (name, phone_number, vehicle_type) VALUES
('Ucup', '081234567893', 'Motorcycle'),
('Udin', '081234567894', 'Car'),
('Ujang', '081234567895', 'Motorcycle');


-- add Locations
INSERT INTO Locations (name, type) VALUES
('Jalan Cilengsi Bogor', 'pickup'),
('Jalan Parung Bogor', 'pickup'),
('Jalan Ciputat Tangerang', 'dropoff'),
('Jalan Pamulang Tangerang', 'dropoff');

-- add Orders
INSERT INTO Orders (customer_id, driver_id, order_date, pickup_location, dropoff_location, total_fare) VALUES
(1, 1, '2024-10-01 08:15:00', 1, 3, 50000),
(1, 2, '2024-10-10 14:00:00', 2, 4, 75000),
(2, 3, '2024-11-05 18:30:00', 1, 4, 30000),
(3, 1, '2024-10-20 09:45:00', 2, 3, 55000),
(3, 2, '2024-11-10 12:00:00', 1, 4, 60000);

-- add Orders Status
INSERT INTO OrderStatus (order_id, status, updated_at) VALUES
(1, 'completed', '2024-10-01 09:00:00'),
(2, 'completed', '2024-10-10 15:00:00'),
(3, 'ongoing', '2024-11-05 18:45:00'),
(4, 'completed', '2024-10-20 10:30:00'),
(5, 'cancelled', '2024-11-10 12:15:00');


SELECT * FROM customers;

ALTER TABLE COLUMN customers;


SELECT * FROM drivers;


SELECT * FROM Orders;

SELECT DATE_TRUNC('month', order_date) AS month, COUNT(order_id) AS total_orders FROM Orders GROUP BY month;

SELECT DATE_TRUNC('month', o.order_date) AS month, o.customer_id, c.name, COUNT(o.order_id) AS total_orders FROM Orders o JOIN customers c ON o.customer_id = c.customer_id GROUP BY month, o.customer_id, c.name ORDER BY month, total_orders DESC;

SELECT L.name AS pickup_location, COUNT(O.order_id) AS total_orders FROM Orders O JOIN Locations L ON O.pickup_location = L.location_id GROUP BY L.name ORDER BY total_orders DESC;

SELECT TO_CHAR(order_date, 'HH24:MI') AS hour_minute, COUNT(order_id) AS total_orders FROM Orders GROUP BY hour_minute ORDER BY total_orders DESC;

SELECT cl.customer_id, c.name,c.phone_number,c.email, COUNT(CASE WHEN cl.logout_time IS NULL THEN 1 END) AS total_logins,COUNT(CASE WHEN cl.logout_time IS NOT NULL THEN 1 END) AS total_logouts FROM CustomerLogins cl JOIN Customers c ON c.customer_id = cl.customer_id GROUP BY cl.customer_id, c.name, c.phone_number, c.email;

SELECT DATE_TRUNC('month', o.order_date) AS month, d.driver_id, d.name, COUNT(CASE WHEN os.status = 'completed' THEN o.order_id END) AS total_success_orders, COUNT(CASE WHEN os.status = 'cancelled' THEN o.order_id END) AS total_cancelled_orders,COUNT(o.order_id) AS total_orders FROM Orders o JOIN Drivers d ON o.driver_id = d.driver_id JOIN OrderStatus os ON os.order_id = o.order_id GROUP BY month, d.driver_id, d.name ORDER BY month, total_orders DESC;