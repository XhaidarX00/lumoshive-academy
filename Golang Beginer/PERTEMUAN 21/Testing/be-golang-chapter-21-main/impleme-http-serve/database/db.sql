-- Active: 1730083286169@@127.0.0.1@5432@webgolang@public
CREATE TABLE table_name(  
    id int NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    create_time DATE,
    name VARCHAR(255)
);
COMMENT ON TABLE table_name IS '';
COMMENT ON COLUMN table_name.name IS '';