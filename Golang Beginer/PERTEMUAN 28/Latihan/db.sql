-- Active: 1730083286169@@127.0.0.1@5432@travel@public
CREATE TABLE place (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    photo_url TEXT NOT NULL
    price INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    transaction_id int,
    rating DECIMAL(2, 1),
    review_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (transaction_id) REFERENCES transaction(id) ON DELETE CASCADE
);


CREATE TABLE event (
    id SERIAL PRIMARY KEY,
    place_id INT,
    date_event TIMESTAMP,
    FOREIGN KEY (place_id) REFERENCES place(id) ON DELETE CASCADE
);


CREATE TABLE transaction (
    id SERIAL PRIMARY KEY,
    event_id INT,
    status_order BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (event_id) REFERENCES event(id) ON DELETE CASCADE
);


SELECT 
    p.id, 
    p.name, 
    p.description,
    p.photo_url,
    p.price,
    e.date_event,
    SUM(CASE WHEN t.status_order = TRUE THEN 1 ELSE 0 END) AS people,
    COALESCE(ROUND(AVG(r.rating), 1), 0) AS rating
FROM event e
JOIN place p ON e.place_id = p.id
JOIN transaction t ON t.event_id = e.id
LEFT JOIN reviews r ON r.transaction_id = t.id
GROUP BY p.id, p.name, p.description, p.price, e.date_event;






INSERT INTO place (name, description, price) VALUES
('Bali Beach', 'Pantai indah di Bali', 150000),
('Mount Bromo', 'Gunung terkenal di Jawa Timur', 200000),
('Raja Ampat Islands', 'Kepulauan eksotis di Papua', 500000),
('Toba Lake', 'Danau terbesar di Indonesia', 100000),
('Komodo Island', 'Habitat asli Komodo', 300000),
('Borobudur Temple', 'Candi terkenal di Jawa Tengah', 250000),
('Nusa Penida', 'Pulau kecil dengan pantai yang indah', 180000),
('Derawan Islands', 'Kepulauan dengan ekosistem laut kaya', 400000),
('Wakatobi', 'Taman Nasional Laut di Sulawesi', 450000),
('Lake Toba', 'Danau vulkanik terbesar di Sumatera', 120000),
('Gili Islands', 'Pulau wisata dekat Lombok', 210000),
('Ijen Crater', 'Kawah dengan api biru di Jawa Timur', 220000),
('Puncak Jaya', 'Puncak tertinggi di Indonesia', 600000),
('Green Canyon', 'Sungai dengan pemandangan spektakuler', 130000),
('Mentawai Islands', 'Destinasi surfing terkenal', 280000),
('Papua Jungle', 'Hutan hujan alami di Papua', 400000),
('Tanjung Puting', 'Taman nasional dengan orangutan', 250000),
('Ubud', 'Pusat budaya dan seni di Bali', 140000);


INSERT INTO event (place_id, date_event) VALUES
(1, '2024-11-01 10:00:00'),
(2, '2024-11-02 11:00:00'),
(3, '2024-11-03 12:00:00'),
(4, '2024-11-04 13:00:00'),
(5, '2024-11-05 14:00:00'),
(6, '2024-11-06 15:00:00'),
(7, '2024-11-07 16:00:00'),
(8, '2024-11-08 17:00:00'),
(9, '2024-11-09 18:00:00'),
(10, '2024-11-10 19:00:00'),
(11, '2024-11-11 20:00:00'),
(12, '2024-11-12 21:00:00'),
(13, '2024-11-13 22:00:00'),
(14, '2024-11-14 23:00:00'),
(15, '2024-11-15 08:00:00'),
(16, '2024-11-16 09:00:00'),
(17, '2024-11-17 10:00:00'),
(18, '2024-11-18 11:00:00');


INSERT INTO transaction (event_id, status_order) VALUES
(1, TRUE),
(2, FALSE),
(3, TRUE),
(4, FALSE),
(5, TRUE),
(6, TRUE),
(7, FALSE),
(8, TRUE),
(9, TRUE);
-- (10, FALSE),
-- (11, TRUE),
-- (12, TRUE),
-- (13, FALSE),
-- (14, TRUE),
-- (15, TRUE),
-- (16, FALSE),
-- (17, TRUE),
-- (18, TRUE);


INSERT INTO reviews (transaction_id, rating) VALUES
(1, 4.5),
(3, 5.0),
(5, 4.7),
(6, 3.8),
(8, 5.0),
(9, 4.1),
(11, 4.3),
(12, 3.7),
(14, 4.4),
(15, 4.8),
(17, 4.0),
(18, 4.9);
