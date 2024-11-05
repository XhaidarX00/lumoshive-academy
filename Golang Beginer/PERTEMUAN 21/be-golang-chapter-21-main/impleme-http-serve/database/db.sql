-- Active: 1730083286169@@127.0.0.1@5432@webgolang@public
CREATE TABLE customers (
    id       VARCHAR(10) PRIMARY KEY,
	username VARCHAR(10) NOT NULL,
	password VARCHAR(20) NOT NULL,
	email    VARCHAR(50) NOT NULL UNIQUE
);

INSERT INTO customers (id, username, password, email) VALUES
('cst1', 'customer1', 'pass', 'customer1@excample.com');

SELECT username, password, email FROM customers WHERE username='customer1' AND password='pass';


CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(10) UNIQUE NOT NULL,
    password VARCHAR(20) NOT NULL,
    role ENUM('admin', 'dev') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tasks (
    task_id SERIAL PRIMARY KEY,
    user_id INT,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status ENUM('pending', 'completed') DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE SET NULL
);

CREATE TABLE tokens (
    token_id SERIAL PRIMARY KEY,
    user_id INT,
    token VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Create ENUM types
CREATE TYPE user_role AS ENUM ('admin', 'dev');
CREATE TYPE task_status AS ENUM ('pending', 'completed');

-- Create users table
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(10) UNIQUE NOT NULL,
    password VARCHAR(20) NOT NULL,
    role user_role NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

SELECT username, password, role FROM users WHERE username='admin1' AND password='password123';


-- Create tasks table
CREATE TABLE tasks (
    task_id SERIAL PRIMARY KEY,
    user_id INT,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status task_status DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE SET NULL
);

-- Create a trigger function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create a trigger for tasks table
CREATE TRIGGER update_task_updated_at
BEFORE UPDATE ON tasks
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();

-- Create tokens table
CREATE TABLE tokens (
    token_id SERIAL PRIMARY KEY,
    user_id INT,
    token VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);


-- Insert dummy data into users table
INSERT INTO users (username, password, role)
VALUES
    ('admin1', 'password123', 'admin'),
    ('dev1', 'password456', 'dev');
    

-- Insert dummy data into tasks table
INSERT INTO tasks (user_id, title, description, status)
VALUES
    (1, 'Setup project repository', 'Create a new Git repository for the project.', 'pending'),
    (1, 'Design database schema', 'Outline the database tables and relationships.', 'completed'),
    (2, 'Implement authentication', 'Set up user login and registration features.', 'pending'),
    (2, 'Create API endpoints', 'Develop the RESTful API for the application.', 'pending'),
    (1, 'Review code', 'Check the code quality and suggest improvements.', 'completed'),
    (2, 'Write tests', 'Create unit tests for the main functionalities.', 'pending');

-- Insert dummy data into tokens table
INSERT INTO tokens (user_id, token, expires_at)
VALUES
    (1, 'token_admin_1', NOW() + INTERVAL '5 minutes'),
    (2, 'token_dev_1', NOW() + INTERVAL '5 minutes');

