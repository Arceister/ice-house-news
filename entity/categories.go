package entity

import "github.com/google/uuid"

type Categories struct {
	Id   uuid.UUID `json:"categories_id"`
	Name *string   `json:"name"`
}
