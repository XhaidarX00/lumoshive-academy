-- Active: 1730083286169@@127.0.0.1@5432@bookstore@public


CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    username VARCHAR(20) NOT NULL UNIQUE,
    password VARCHAR(20) NOT NULL,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(15),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tokens (
    token_id SERIAL PRIMARY KEY,
    customer_id INT,
    token VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
);

DROP TABLE customers;


CREATE TABLE address (
    id SERIAL PRIMARY KEY,
    street VARCHAR(255),
    city VARCHAR(100),
    postal VARCHAR(10),
    country VARCHAR(100),
    FOREIGN KEY (id) REFERENCES customers(id) ON DELETE CASCADE
);

CREATE TABLE books (
    id VARCHAR(10) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(100) NOT NULL,
    author VARCHAR(100) NOT NULL,
    price INT NOT NULL,
    discount DECIMAL(5, 2)
);


CREATE TYPE order_status AS ENUM ('Pending', 'Processing', 'Completed');

SELECT id, customer_id, payment_methode, total_amount, discount, final_amount, order_date, status FROM orders ORDER BY id;

SELECT u.name
FROM customers u
JOIN orders o ON o.customer_id = u.id 
WHERE o.customer_id = 1
LIMIT 1;


CREATE TABLE orders (
    id VARCHAR(20) PRIMARY KEY,
    customer_id INT,
    payment_methode VARCHAR(50),
    total_amount INT,
    discount DECIMAL(5, 2),
    final_amount INT,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status order_status NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
);

CREATE TABLE order_items (
    id VARCHAR(10) PRIMARY KEY,
    order_id VARCHAR(20),
    book_id VARCHAR(20),
    quantity INT NOT NUll CHECK (quantity > 0),
    subtotal INT,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
);

CREATE TABLE reviews (
    id VARCHAR(20) PRIMARY KEY,
    order_id VARCHAR(20),
    book_id VARCHAR(20),
    customer_id INT,
    rating DECIMAL(2, 1),
    review_text TEXT,
    review_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
);

DROP TABLE reviews;


INSERT INTO customers(
    username, password, name, phone_number
) VALUES('customer1', 'pass', 'Haidar', '089230983');

INSERT INTO tokens (customer_id, token, expires_at)
VALUES
    (1, 'token_admin_1', NOW() + INTERVAL '30 day');

INSERT INTO address(
    street, city, postal, country
) VALUES('Jl. Raya Pamijahan', 'Bogor', '16450', 'Indonesia');


-- Menambahkan data dummy ke dalam tabel books
INSERT INTO books (id, name, type, author, price, discount) VALUES
('BK1', 'The Great Gatsby', 'Novel', 'F. Scott Fitzgerald', 100000, 10.00),
('BK2', '1984', 'Dystopian', 'George Orwell', 85000, 15.50),
('BK3', 'To Kill a Mockingbird', 'Novel', 'Harper Lee', 95000, 5.75),
('BK4', 'The Catcher in the Rye', 'Classic', 'J.D. Salinger', 120000, 20.00),
('BK5', 'Pride and Prejudice', 'Romance', 'Jane Austen', 70000, 0.00),
('BK6', 'The Hobbit', 'Fantasy', 'J.R.R. Tolkien', 110000, 25.50),
('BK7', 'Moby Dick', 'Adventure', 'Herman Melville', 135000, 18.00),
('BK8', 'War and Peace', 'Historical', 'Leo Tolstoy', 150000, 5.00),
('BK9', 'Brave New World', 'Dystopian', 'Aldous Huxley', 95000, 12.50),
('BK10', 'The Lord of the Rings', 'Fantasy', 'J.R.R. Tolkien', 200000, 30.00);


INSERT INTO orders (id, customer_id, payment_methode, total_amount, discount, final_amount, order_date, status)
VALUES ('ORD1', 1, 'credit_card', 150000, 10, 135000, '2024-11-07 10:00:00', 'Processing');

INSERT INTO order_items (id, order_id, book_id, quantity, subtotal)
VALUES 
('ORDITMS1', 'ORD1', 'BK1', 2, 100000),
('ORDITMS2', 'ORD1', 'BK1', 1, 50000);

TRUNCATE order_items;


SELECT  c.id, c.name, c.phone_number, a.street, a.city, a.country, a.postal
FROM customers c
JOIN address a ON c.id = a.id;


SELECT o.id, c.name
FROM orders o
JOIN order_items oi ON oi.
JOIN customers c ON c.id = o.customer_id
WHERE o.id = 'ORD1';

SELECT o.id, c.name, oi.book_id, b.name, oi.quantity, b.price, oi.subtotal, 
       o.total_amount, o.discount, o.final_amount, o.order_date, o.status
FROM orders o
JOIN customers c ON o.customer_id = c.id
JOIN order_items oi ON o.id = oi.order_id
JOIN books b ON oi.book_id = b.id
WHERE o.id = 'ORD1';


INSERT INTO reviews (id, order_id, book_id, customer_id, rating, review_text, review_date)
VALUES 
('REV1', 'ORD1', 'BK1', 1, 4.5, 'Buku yang sangat informatif dan mudah dipahami.', '2024-11-07 15:30:00');

TRUNCATE reviews;

SELECT r.id, r.order_id, r.book_id, c.name, r.rating, r.review_text, r.review_date
FROM reviews r
JOIN customers c ON r.customer_id = c.id
WHERE r.order_id = 'ORD1' AND r.book_id = 'BK1';


SELECT MAX(rating) AS highest_rating FROM reviews;

SELECT COUNT(id) FROM orders;

SELECT COUNT(id) FROM books;




SELECT * FROM orders;
SELECT * FROM books;
SELECT * FROM reviews;


INSERT INTO customers (username, password, name, phone_number) VALUES
('johndoe', 'password123', 'John Doe', '081234567890'),
('janedoe', 'password456', 'Jane Doe', '082345678901'),
('alexsmith', 'password789', 'Alex Smith', '083456789012'),
('lucyjones', 'password101', 'Lucy Jones', '084567890123'),
('markbrown', 'password102', 'Mark Brown', '085678901234'),
('emilydavis', 'password103', 'Emily Davis', '086789012345'),
('chrisjohnson', 'password104', 'Chris Johnson', '087890123456'),
('patriciamartin', 'password105', 'Patricia Martin', '088901234567'),
('danielwilliams', 'password106', 'Daniel Williams', '089012345678'),
('susanwhite', 'password107', 'Susan White', '090123456789');

INSERT INTO tokens (customer_id, token, expires_at) VALUES
(1, 'token12345abcd', '2024-12-31 23:59:59'),
(2, 'token67890efgh', '2024-12-31 23:59:59'),
(3, 'token11223ijkl', '2024-12-31 23:59:59'),
(4, 'token44556mnop', '2024-12-31 23:59:59'),
(5, 'token78901qrst', '2024-12-31 23:59:59'),
(6, 'token33445uvwx', '2024-12-31 23:59:59'),
(7, 'token55667yzab', '2024-12-31 23:59:59'),
(8, 'token77889cdef', '2024-12-31 23:59:59'),
(9, 'token99001ghij', '2024-12-31 23:59:59'),
(10, 'token22334klmn', '2024-12-31 23:59:59');

INSERT INTO address (id, street, city, postal, country) VALUES
(1, '123 Main St', 'New York', '10001', 'USA'),
(2, '456 Oak Ave', 'Los Angeles', '90001', 'USA'),
(3, '789 Pine Rd', 'Chicago', '60601', 'USA'),
(4, '101 Maple Blvd', 'Houston', '77001', 'USA'),
(5, '202 Birch Dr', 'Phoenix', '85001', 'USA'),
(6, '303 Cedar Ln', 'Philadelphia', '19101', 'USA'),
(7, '404 Elm St', 'San Antonio', '78201', 'USA'),
(8, '505 Walnut St', 'San Diego', '92101', 'USA'),
(9, '606 Cherry St', 'Dallas', '75201', 'USA'),
(10, '707 Redwood St', 'Austin', '73301', 'USA');

INSERT INTO books (id, name, type, author, price, discount) VALUES
('B001', 'The Great Gatsby', 'Novel', 'F. Scott Fitzgerald', 100000, 10.00),
('B002', '1984', 'Dystopian', 'George Orwell', 85000, 15.50),
('B003', 'To Kill a Mockingbird', 'Novel', 'Harper Lee', 95000, 5.75),
('B004', 'The Catcher in the Rye', 'Classic', 'J.D. Salinger', 120000, 20.00),
('B005', 'Pride and Prejudice', 'Romance', 'Jane Austen', 70000, 0.00),
('B006', 'The Hobbit', 'Fantasy', 'J.R.R. Tolkien', 110000, 25.50),
('B007', 'Moby Dick', 'Adventure', 'Herman Melville', 135000, 18.00),
('B008', 'War and Peace', 'Historical', 'Leo Tolstoy', 150000, 5.00),
('B009', 'Brave New World', 'Dystopian', 'Aldous Huxley', 95000, 12.50),
('B010', 'The Lord of the Rings', 'Fantasy', 'J.R.R. Tolkien', 200000, 30.00);

INSERT INTO orders (id, customer_id, payment_methode, total_amount, discount, final_amount, status) VALUES
('O001', 1, 'Credit Card', 200000, 10.00, 180000, 'Completed'),
('O002', 2, 'PayPal', 150000, 5.00, 142500, 'Completed'),
('O003', 3, 'Debit Card', 120000, 20.00, 96000, 'Completed'),
('O004', 4, 'Bank Transfer', 220000, 0.00, 220000, 'Pending'),
('O005', 5, 'Credit Card', 180000, 15.00, 153000, 'Processing'),
('O006', 6, 'Cash on Delivery', 250000, 10.00, 225000, 'Completed'),
('O007', 7, 'PayPal', 300000, 25.00, 225000, 'Completed'),
('O008', 8, 'Debit Card', 280000, 5.00, 266000, 'Completed'),
('O009', 9, 'Bank Transfer', 100000, 0.00, 100000, 'Pending'),
('O010', 10, 'Credit Card', 350000, 30.00, 245000, 'Completed');

INSERT INTO order_items (id, order_id, book_id, quantity, subtotal) VALUES
('I001', 'O001', 'B001', 2, 200000),
('I002', 'O002', 'B002', 1, 85000),
('I003', 'O003', 'B003', 1, 95000),
('I004', 'O004', 'B004', 3, 360000),
('I005', 'O005', 'B005', 2, 140000),
('I006', 'O006', 'B006', 2, 220000),
('I007', 'O007', 'B007', 1, 135000),
('I008', 'O008', 'B008', 1, 150000),
('I009', 'O009', 'B009', 2, 190000),
('I010', 'O010', 'B010', 1, 200000);

INSERT INTO reviews (id, order_id, book_id, customer_id, rating, review_text) VALUES
('R001', 'O001', 'B001', 1, 4.5, 'Great read! Very engaging and thoughtful.'),
('R002', 'O002', 'B002', 2, 4.0, 'Dystopian future that is still relevant today.'),
('R003', 'O003', 'B003', 3, 5.0, 'A classic, timeless story. Highly recommend!'),
('R004', 'O004', 'B004', 4, 3.5, 'Interesting but a bit slow in the middle.'),
('R005', 'O005', 'B005', 5, 4.8, 'Love the romance and witty dialogue.'),
('R006', 'O006', 'B006', 6, 5.0, 'A wonderful adventure that never gets old.'),
('R007', 'O007', 'B007', 7, 4.2, 'A good adventure story with great character depth.'),
('R008', 'O008', 'B008', 8, 3.8, 'A bit long, but the historical context is valuable.'),
('R009', 'O009', 'B009', 9, 4.5, 'A thought-provoking book that is still relevant.'),
('R010', 'O010', 'B010', 10, 5.0, 'Epic! A fantastic journey from start to finish.');

count(idproduck) around

SELECT * FROM customers;

SELECT u.id, u.name, u.phone_number 
FROM customers u
JOIN orders o ON o.customer_id = u.id 
WHERE o.customer_id = 1
LIMIT;


CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    photo VARCHAR(255),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);


INSERT INTO payments (name, photo, is_active) VALUES 
('Credit Card', 'https://example.com/credit-card.png', true),
('Ewallet', 'https://example.com/paypal.png', true),
('Bank Transfer', 'https://example.com/bank-transfer.png', false);


SELECT * FROM orders;

SELECT 
	id, 
	customer_id, 
	payment_methode, 
	total_amount, 
	discount, 
	final_amount, 
	order_date, 
	status 
	FROM orders 
    WHERE id = 'O001'
    ORDER BY id;
    

SELECT * FROM reviews;

SELECT ROUND(AVG(rating), 1)
			FROM (
				SELECT rating 
				FROM reviews 
				WHERE book_id = 'B002'
				ORDER BY review_date DESC 
				LIMIT 25
			) AS latest_reviews;


SELECT ROUND(AVG(rating), 1) FROM reviews WHERE book_id ='B002';

SELECT * FROM orders;