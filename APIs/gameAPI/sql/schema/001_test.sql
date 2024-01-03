-- +goose Up

CREATE TABLE test (
    name VARCHAR(55) NOT NULL
);

-- +goose Down
DROP TABLE test;
