// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package article

import (
	"github.com/google/uuid"
)

type Article struct {
	ID       uuid.UUID
	AuthorID uuid.UUID
	Title    string
	Content  string
}
