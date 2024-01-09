<<<<<<< HEAD
-- +goose Up

CREATE TABLE catalog_games (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL, 
    description VARCHAR(100),
    imageURL TEXT 
);

-- +goose Down
DROP TABLE catalog_games;
=======
-- +goose Up

CREATE TABLE catalog_games (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL, 
    description VARCHAR(100),
    imageURL TEXT 
);

-- +goose Down
DROP TABLE catalog_games;
>>>>>>> 010c1db92ac649cd1bf4f35f1618dea9c537bbf2
