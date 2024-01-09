-- +goose Up
CREATE TABLE tasks (
    id UUID PRIMARY KEY,  
    description VARCHAR(100), 
    game_id UUID REFERENCES catalog_games(id) NOT NULL
);

-- +goose Down
DROP TABLE tasks;
