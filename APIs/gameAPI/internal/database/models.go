// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type CatalogGame struct {
	ID          uuid.UUID
	Name        string
	Description sql.NullString
	Imageurl    sql.NullString
}

type Task struct {
	ID          uuid.UUID
	Description sql.NullString
	GameID      uuid.UUID
}
