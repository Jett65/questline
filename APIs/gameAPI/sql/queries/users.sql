-- name: CreateUser :one
INSERT INTO users (id, username, email, password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users WHERE id=$1;

-- name: UpdateUser :one
UPDATE users SET username=$2, email=$3, password=$4
WHERE id=$1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id=$1;

