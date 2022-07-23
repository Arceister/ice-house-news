package entity

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	Id          uuid.UUID `json:"id"`
	NewsId      uuid.UUID `json:"news_id"`
	UsersId     uuid.UUID `json:"users_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type CommentInsertRequest struct {
	Description string `json:"description" validate:"required"`
}
