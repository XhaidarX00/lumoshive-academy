-- Active: 1730083286169@@127.0.0.1@5432@travel@public
CREATE TABLE place (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    photo_url TEXT NOT NULL,
    price INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    transaction_id INT,
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
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    message TEXT,
    event_id INT,
    status_order BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (event_id) REFERENCES event(id) ON DELETE CASCADE
);


CREATE TABLE gallery (
    id SERIAL PRIMARY KEY,
    photo_url TEXT NOT NULL,
    place_id INT,
    description TEXT NOT NULL
);




SELECT 
    p.id, 
    e.id AS event_id,
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
WHERE DATE(e.date_event) = '2024-11-08'
GROUP BY p.id, e.id, p.name, p.description, p.price, e.date_event
ORDER BY p.id
LIMIT 6 OFFSET 1;


SELECT 
    p.id, 
    e.id AS event_id,
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
JOIN reviews r ON r.transaction_id = t.id
WHERE DATE(e.date_event) = '2024-12-02'
GROUP BY p.id, e.id, p.name, p.description, p.price, e.date_event
ORDER BY e.id;


SELECT
    p.id,
    e.id AS event_id,
    p.name,
    p.description,
    p.price,
    e.date_event,
    SUM(CASE WHEN t.status_order = TRUE THEN 1 ELSE 0 END) AS people,
    COALESCE(ROUND(AVG(r.rating), 1), 0) AS rating
FROM event e
JOIN place p ON e.place_id = p.id
JOIN transaction t ON t.event_id = e.id
JOIN reviews r ON r.transaction_id = t.id 
WHERE DATE(e.date_event) = '2024-11-06'
GROUP BY p.id, e.id, p.name, p.description, p.price, e.date_event ORDER BY e.id;


SELECT * FROM event;

SELECT 
    p.id, 
    p.name, 
    p.description,
    p.price,
    e.date_event,
    SUM(CASE WHEN t.status_order = TRUE THEN 1 ELSE 0 END) AS people,
    COUNT(r.rating) AS rating
FROM event e
JOIN place p ON e.place_id = p.id
JOIN transaction t ON t.event_id = e.id
LEFT JOIN reviews r ON r.transaction_id = t.id
WHERE p.id = 1
GROUP BY p.id, p.name, p.description, p.price, e.date_event
ORDER BY p.id;

SELECT 
    g.id, 
    g.photo_url,
    g.description
FROM gallery g
JOIN place p ON g.place_id = p.id
WHERE p.id = 1
GROUP BY g.id, g.photo_url, g.description
ORDER BY g.id;


INSERT INTO transaction (name, email, phone_number, message, event_id, status_order)
VALUES
('Ali', 'ali@example.com', '081234567890', 'Booking untuk liburan keluarga', 1, TRUE);


SELECT * FROM transaction;

-- Insert data dummy untuk tabel `place`
INSERT INTO place (name, description, photo_url, price)
VALUES 
('Bali Beach', 'Pantai dengan pemandangan indah dan pasir putih.', 'https://example.com/photo_bali.jpg', 1500000),
('Jakarta Museum', 'Museum sejarah dengan banyak peninggalan bersejarah.', 'https://example.com/photo_jakarta.jpg', 500000),
('Bandung Park', 'Taman kota dengan suasana sejuk dan alami.', 'https://example.com/photo_bandung.jpg', 250000),
('Lombok Hill', 'Bukit dengan pemandangan sunset yang memukau.', 'https://example.com/photo_lombok.jpg', 1000000),
('Yogyakarta Temple', 'Candi peninggalan kerajaan kuno.', 'https://example.com/photo_yogyakarta.jpg', 800000),
('Surabaya Zoo', 'Kebun binatang terbesar di Jawa Timur.', 'https://example.com/photo_surabaya.jpg', 400000),
('Makassar Beach', 'Pantai dengan pasir hitam yang unik.', 'https://example.com/photo_makassar.jpg', 750000),
('Medan Lake', 'Danau yang tenang dan damai.', 'https://example.com/photo_medan.jpg', 550000),
('Bromo Mountain', 'Gunung dengan kawah yang megah.', 'https://example.com/photo_bromo.jpg', 1200000),
('Raja Ampat Island', 'Pulau dengan keindahan laut dan karang.', 'https://example.com/photo_raja_ampat.jpg', 2000000),
('Manado Diving', 'Spot diving terbaik dengan keanekaragaman laut.', 'https://example.com/photo_manado.jpg', 1800000),
('Malang Garden', 'Kebun yang penuh dengan bunga-bunga indah.', 'https://example.com/photo_malang.jpg', 600000);

-- Insert data dummy untuk tabel `reviews`
INSERT INTO reviews (transaction_id, rating)
VALUES
(1, 4.5),
(2, 3.8),
(3, 4.2),
(4, 5.0),
(5, 4.0),
(6, 3.5),
(7, 4.7),
(8, 4.1),
(9, 3.9),
(10, 4.6),
(11, 4.3),
(12, 5.0);

-- Insert data dummy untuk tabel `event`
INSERT INTO event (place_id, date_event)
VALUES
(1, '2024-12-01 10:00:00'),
(2, '2024-12-02 11:00:00'),
(3, '2024-12-03 14:00:00'),
(4, '2024-12-04 13:00:00'),
(5, '2024-12-05 15:00:00'),
(6, '2024-12-06 16:00:00'),
(7, '2024-12-07 09:00:00'),
(8, '2024-12-08 10:00:00'),
(9, '2024-12-09 11:00:00'),
(10, '2024-12-10 14:00:00'),
(11, '2024-12-11 15:00:00'),
(12, '2024-12-12 16:00:00');

-- Insert data dummy untuk tabel `transaction`
INSERT INTO transaction (name, email, phone_number, message, event_id, status_order)
VALUES
('Ali', 'ali@example.com', '081234567890', 'Booking untuk liburan keluarga', 1, TRUE),
('Budi', 'budi@example.com', '081234567891', 'Mau ambil foto prewedding', 2, TRUE),
('Citra', 'citra@example.com', '081234567892', 'Liburan akhir tahun', 3, FALSE),
('Dewi', 'dewi@example.com', '081234567893', 'Event gathering kantor', 4, TRUE),
('Eko', 'eko@example.com', '081234567894', 'Family trip', 5, TRUE),
('Fajar', 'fajar@example.com', '081234567895', 'Booking untuk bulan depan', 6, FALSE),
('Gita', 'gita@example.com', '081234567896', 'Liburan akhir pekan', 7, TRUE),
('Hana', 'hana@example.com', '081234567897', 'Honeymoon', 8, TRUE),
('Irma', 'irma@example.com', '081234567898', 'Liburan bersama anak-anak', 9, FALSE),
('Joko', 'joko@example.com', '081234567899', 'Trip bisnis', 10, TRUE),
('Kevin', 'kevin@example.com', '081234567800', 'Liburan dengan teman-teman', 11, TRUE),
('Lia', 'lia@example.com', '081234567801', 'Family gathering', 12, TRUE);


-- Insert data dummy untuk tabel `gallery` dengan 6 data untuk setiap `place`
-- Data dummy untuk `gallery` setiap `place`
INSERT INTO gallery (photo_url, place_id, description)
VALUES
-- Bali Beach
('https://example.com/photo_bali1.jpg', 1, 'Pemandangan laut Bali di sore hari'),
('https://example.com/photo_bali2.jpg', 1, 'Peselancar di pantai Bali'),
('https://example.com/photo_bali3.jpg', 1, 'Matahari terbenam di Bali'),
('https://example.com/photo_bali4.jpg', 1, 'Pasir putih Bali'),
('https://example.com/photo_bali5.jpg', 1, 'Pohon kelapa di pantai Bali'),
('https://example.com/photo_bali6.jpg', 1, 'Perahu nelayan Bali'),
-- Jakarta Museum
('https://example.com/photo_jakarta1.jpg', 2, 'Arsitektur klasik museum'),
('https://example.com/photo_jakarta2.jpg', 2, 'Artefak bersejarah di Jakarta'),
('https://example.com/photo_jakarta3.jpg', 2, 'Lukisan di museum'),
('https://example.com/photo_jakarta4.jpg', 2, 'Ruang pameran Jakarta Museum'),
('https://example.com/photo_jakarta5.jpg', 2, 'Patung-patung di museum'),
('https://example.com/photo_jakarta6.jpg', 2, 'Taman museum Jakarta'),
-- Bandung Park
('https://example.com/photo_bandung1.jpg', 3, 'Taman dengan banyak pohon rindang'),
('https://example.com/photo_bandung2.jpg', 3, 'Danau di tengah taman Bandung'),
('https://example.com/photo_bandung3.jpg', 3, 'Suasana sejuk di taman Bandung'),
('https://example.com/photo_bandung4.jpg', 3, 'Jalan setapak di Bandung Park'),
('https://example.com/photo_bandung5.jpg', 3, 'Piknik di taman Bandung'),
('https://example.com/photo_bandung6.jpg', 3, 'Bunga warna-warni di Bandung Park'),
-- Lombok Hill
('https://example.com/photo_lombok1.jpg', 4, 'Pemandangan dari puncak Lombok Hill'),
('https://example.com/photo_lombok2.jpg', 4, 'Sunset di Lombok Hill'),
('https://example.com/photo_lombok3.jpg', 4, 'Rute pendakian Lombok Hill'),
('https://example.com/photo_lombok4.jpg', 4, 'Pohon-pohon rindang di Lombok Hill'),
('https://example.com/photo_lombok5.jpg', 4, 'Pemandangan laut dari Lombok Hill'),
('https://example.com/photo_lombok6.jpg', 4, 'Jalur trekking di Lombok Hill'),
-- Yogyakarta Temple
('https://example.com/photo_yogyakarta1.jpg', 5, 'Candi Borobudur di Yogyakarta'),
('https://example.com/photo_yogyakarta2.jpg', 5, 'Patung di sekitar candi Yogyakarta'),
('https://example.com/photo_yogyakarta3.jpg', 5, 'Sunrise di candi Yogyakarta'),
('https://example.com/photo_yogyakarta4.jpg', 5, 'Situs candi yang indah di Yogyakarta'),
('https://example.com/photo_yogyakarta5.jpg', 5, 'Pemandangan candi dari kejauhan'),
('https://example.com/photo_yogyakarta6.jpg', 5, 'Ukiran di dinding candi Yogyakarta'),
-- Surabaya Zoo
('https://example.com/photo_surabaya1.jpg', 6, 'Hewan langka di Surabaya Zoo'),
('https://example.com/photo_surabaya2.jpg', 6, 'Kandang gajah di Surabaya Zoo'),
('https://example.com/photo_surabaya3.jpg', 6, 'Taman bermain di kebun binatang Surabaya'),
('https://example.com/photo_surabaya4.jpg', 6, 'Pengunjung melihat singa di Surabaya Zoo'),
('https://example.com/photo_surabaya5.jpg', 6, 'Koleksi burung langka di Surabaya Zoo'),
('https://example.com/photo_surabaya6.jpg', 6, 'Pemberian makan hewan di Surabaya Zoo'),
-- Makassar Beach
('https://example.com/photo_makassar1.jpg', 7, 'Pantai Makassar dengan ombak besar'),
('https://example.com/photo_makassar2.jpg', 7, 'Pasir hitam di pantai Makassar'),
('https://example.com/photo_makassar3.jpg', 7, 'Aktivitas selancar di Makassar Beach'),
('https://example.com/photo_makassar4.jpg', 7, 'Perahu nelayan di Makassar Beach'),
('https://example.com/photo_makassar5.jpg', 7, 'Keindahan matahari terbenam di Makassar Beach'),
('https://example.com/photo_makassar6.jpg', 7, 'Kegiatan pantai di Makassar'),
-- Medan Lake
('https://example.com/photo_medan1.jpg', 8, 'Danau Medan di pagi hari'),
('https://example.com/photo_medan2.jpg', 8, 'Suasana danau yang tenang di Medan'),
('https://example.com/photo_medan3.jpg', 8, 'Perahu di Danau Medan'),
('https://example.com/photo_medan4.jpg', 8, 'Tepi danau Medan dengan pepohonan'),
('https://example.com/photo_medan5.jpg', 8, 'Aktivitas perahu di Danau Medan'),
('https://example.com/photo_medan6.jpg', 8, 'Pemandangan indah danau Medan'),
-- Bromo Mountain
('https://example.com/photo_bromo1.jpg', 9, 'Pemandangan kawah Bromo'),
('https://example.com/photo_bromo2.jpg', 9, 'Gunung Bromo dengan kabut tipis'),
('https://example.com/photo_bromo3.jpg', 9, 'Sunrise dari puncak Bromo'),
('https://example.com/photo_bromo4.jpg', 9, 'Pasir berbisik di Bromo'),
('https://example.com/photo_bromo5.jpg', 9, 'Jalur pendakian menuju puncak Bromo'),
('https://example.com/photo_bromo6.jpg', 9, 'Keindahan Bromo di malam hari'),
-- Raja Ampat Island
('https://example.com/photo_raja_ampat1.jpg', 10, 'Laut biru jernih di Raja Ampat'),
('https://example.com/photo_raja_ampat2.jpg', 10, 'Keindahan terumbu karang di Raja Ampat'),
('https://example.com/photo_raja_ampat3.jpg', 10, 'Pantai putih di Raja Ampat'),
('https://example.com/photo_raja_ampat4.jpg', 10, 'Menyelam di Raja Ampat'),
('https://example.com/photo_raja_ampat5.jpg', 10, 'Pulau-pulau kecil di Raja Ampat'),
('https://example.com/photo_raja_ampat6.jpg', 10, 'Keindahan bawah laut Raja Ampat'),
-- Manado Diving
('https://example.com/photo_manado1.jpg', 11, 'Pemandangan bawah laut Manado'),
('https://example.com/photo_manado2.jpg', 11, 'Diving di Manado'),
('https://example.com/photo_manado3.jpg', 11, 'Terumbu karang di Manado'),
('https://example.com/photo_manado4.jpg', 11, 'Aktivitas menyelam di Manado'),
('https://example.com/photo_manado5.jpg', 11, 'Keanekaragaman ikan di Manado'),
('https://example.com/photo_manado6.jpg', 11, 'Pemandangan laut Manado yang menakjubkan'),
-- Malang Garden
('https://example.com/photo_malang1.jpg', 12, 'Kebun bunga di Malang'),
('https://example.com/photo_malang2.jpg', 12, 'Suasana sejuk di Malang Garden'),
('https://example.com/photo_malang3.jpg', 12, 'Bunga berwarna-warni di Malang'),
('https://example.com/photo_malang4.jpg', 12, 'Taman dengan pemandangan pegunungan di Malang'),
('https://example.com/photo_malang5.jpg', 12, 'Jalan setapak di Malang Garden'),
('https://example.com/photo_malang6.jpg', 12, 'Pemandangan hijau di Malang Garden');

