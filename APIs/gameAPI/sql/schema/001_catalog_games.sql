-- +goose Up

CREATE TABLE catalog_games (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL, 
    description VARCHAR(100),
    imageURL TEXT 
);

-- +goose Down
DROP TABLE catalog_games;