CREATE TABLE partner (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(50),
    address TEXT,
    logo TEXT,
    aproved BOOLEAN DEFAULT FALSE,
    default_user INTEGER REFERENCES users(id) ON DELETE SET NULL
);