package entity

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	Id          uuid.UUID   `json:"id"`
	Description string      `json:"description"`
	User        Commentator `json:"commentator"`
	CreatedAt   time.Time   `json:"created_at"`
}

type Commentator struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Picture string    `json:"picture"`
}

type CommentInsert struct {
	Id     uuid.UUID
	UserId uuid.UUID
	NewsId uuid.UUID
	CommentInsertRequest
}

type CommentInsertRequest struct {
	Description string `json:"description" validate:"required"`
}
