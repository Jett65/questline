-- name: CreateTask :one
INSERT INTO tasks (id, description, game_id) 
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTask :many
SELECT * FROM tasks;

-- name: GetTaskById :one
SELECT * FROM tasks WHERE id=$1;

-- name: UpdateTask :one
UPDATE tasks SET description=$2, game_id=$3
WHERE id=$1
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id=$1;
