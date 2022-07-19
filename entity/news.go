package entity

import (
	"time"

	"github.com/google/uuid"
)

type NewsListOutput struct {
	Id               uuid.UUID    `json:"id"`
	Title            string       `json:"title"`
	SlugUrl          string       `json:"slug_url"`
	CoverImage       *string      `json:"cover_image"`
	AdditionalImages []string     `json:"additional_images"`
	CreatedAt        time.Time    `json:"created_at"`
	Nsfw             bool         `json:"nsfw"`
	Category         NewsCategory `json:"category"`
	Counter          NewsCounter  `json:"counter"`
}

type NewsCategory struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type NewsCounter struct {
	Upvote   int32 `json:"upvote"`
	Downvote int32 `json:"downvote"`
	Comment  int32 `json:"comment"`
	View     int32 `json:"view"`
}
