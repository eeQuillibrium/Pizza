CREATE TABLE IF NOT EXISTS users 
(
    id SERIAL PRIMARY KEY,
    login VARCHAR(256) NOT NULL,
    passhash VARCHAR(256) NOT NULL
);