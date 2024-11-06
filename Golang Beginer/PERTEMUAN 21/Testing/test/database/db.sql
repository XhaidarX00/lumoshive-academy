-- Active: 1730083286169@@127.0.0.1@5432@webgolang@public

CREATE TABLE users (
    id VARCHAR(10) PRIMARY KEY,
    username VARCHAR(10) NOT NULL,
    password VARCHAR(20) NOT NULL,
);


INSERT INTO users (id, username, password) VALUES
('usr1', 'admin', 'admin123');

SELECT * FROM users;

SELECT id, username, password FROM users WHERE username = 'admin';