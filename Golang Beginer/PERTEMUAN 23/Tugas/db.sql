-- Active: 1730083286169@@127.0.0.1@5432@webgolang@public

CREATE TYPE user_status AS ENUM ('Aktif', 'Tidak Aktif');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    username VARCHAR(10) NOT NULL UNIQUE,
    password VARCHAR(20) NOT NULL CHECK (LENGTH(password) >= 8),
    active user_status NOT NULL DEFAULT 'Aktif',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);


CREATE TABLE tokens (
    token_id SERIAL PRIMARY KEY,
    user_id INT,
    token VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE tasks (
    task_id SERIAL PRIMARY KEY,
    user_id INT,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status task_status DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
);

TRUNCATE users;

INSERT INTO users (name, username, password, active)
VALUES 
('Haidar1', 'admin1', 'password123', 'Tidak Aktif'),
('Haidar2', 'admin2', 'password345', 'Tidak Aktif');

-- Insert dummy data into tasks table
INSERT INTO tasks (user_id, title, description, status)
VALUES
    (1, 'Setup project repository', 'Create a new Git repository for the project.', 'pending'),
    (1, 'Design database schema', 'Outline the database tables and relationships.', 'completed');

-- Insert dummy data into tokens table
INSERT INTO tokens (user_id, token, expires_at)
VALUES
    (1, 'token_admin_1', NOW() + INTERVAL '30 day');

SELECT * FROM tasks;

SELECT * FROM users;

SELECT u.id, u.name, u.username, u.password, t.token, u.active 
FROM users u
JOIN tokens t ON t.user_id = u.id
WHERE u.username = 'admin1' AND u.password = 'password123';