-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL, 
    username VARCHAR(50) NOT NULL, 
    email TEXT NOT NULL, 
    password TEXT NOT NULL 
);

-- +goose Down
DROP TABLE users;
