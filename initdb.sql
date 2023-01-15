-- Connect to the database
\c dbname

-- Create a table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
);

-- Insert data into the table
INSERT INTO users (name, email) VALUES ('John Doe', 'johndoe@example.com');
INSERT INTO users (name, email) VALUES ('Jane Smith', 'janesmith@example.com');
INSERT INTO users (name, email) VALUES ('Bob Johnson', 'bobjohnson@example.com');

-- Select data from the table
SELECT * FROM users;