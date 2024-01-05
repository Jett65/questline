-- +goose Up

CREATE TABLE catalog_games (
    id UUID PRIMARY KEY,
    name VARCHAR, 
    description VARCHAR(100),
    image TEXT
);

-- +goose Down
DROP TABLE catalog_games;
