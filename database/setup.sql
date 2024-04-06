CREATE TABLE IF NOT EXISTS players (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- populate players
INSERT INTO players (name, email, password) VALUES
('Alice', ' [email protected]', 'password'),
('Bob', ' [email protected]', 'password'),
('Charlie', ' [email protected]', 'password');
