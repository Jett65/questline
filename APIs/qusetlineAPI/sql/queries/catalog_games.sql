-- name: CreateCatalogGame :one
INSERT INTO catalog_games (id, name, description, imageURL)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAllCatalogGames :many
SELECT * FROM catalog_games;

-- name: GetCatalogGameById :one
SELECT * FROM catalog_games WHERE id=$1;

-- name: UpdateCatalogGame :one
UPDATE catalog_games SET name=$2, description=$3, imageURL=$4
WHERE id=$1
RETURNING *;

-- name: DeleteCatalogGame :exec
DELETE FROM catalog_games WHERE id=$1;
