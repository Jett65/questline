// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: tasks.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (id, description, game_id) 
VALUES ($1, $2, $3)
RETURNING id, description, game_id
`

type CreateTaskParams struct {
	ID          uuid.UUID
	Description sql.NullString
	GameID      uuid.UUID
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask, arg.ID, arg.Description, arg.GameID)
	var i Task
	err := row.Scan(&i.ID, &i.Description, &i.GameID)
	return i, err
}

const deleteTask = `-- name: DeleteTask :exec
DELETE FROM tasks WHERE id=$1
`

func (q *Queries) DeleteTask(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTask, id)
	return err
}

const getTask = `-- name: GetTask :many
SELECT id, description, game_id FROM tasks
`

func (q *Queries) GetTask(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getTask)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(&i.ID, &i.Description, &i.GameID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTaskByGameId = `-- name: GetTaskByGameId :many
SELECT id, description, game_id FROM tasks WHERE game_id=$1
`

func (q *Queries) GetTaskByGameId(ctx context.Context, gameID uuid.UUID) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getTaskByGameId, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(&i.ID, &i.Description, &i.GameID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTaskById = `-- name: GetTaskById :one
SELECT id, description, game_id FROM tasks WHERE id=$1
`

func (q *Queries) GetTaskById(ctx context.Context, id uuid.UUID) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskById, id)
	var i Task
	err := row.Scan(&i.ID, &i.Description, &i.GameID)
	return i, err
}

const updateTask = `-- name: UpdateTask :one
UPDATE tasks SET description=$2, game_id=$3
WHERE id=$1
RETURNING id, description, game_id
`

type UpdateTaskParams struct {
	ID          uuid.UUID
	Description sql.NullString
	GameID      uuid.UUID
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTask, arg.ID, arg.Description, arg.GameID)
	var i Task
	err := row.Scan(&i.ID, &i.Description, &i.GameID)
	return i, err
}