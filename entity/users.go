package entity

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Bio      *string   `json:"bio"`
	Web      *string   `json:"web"`
	Picture  *string   `json:"picture"`
}

type UserSignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserToken struct {
	Token string `json:"token"`
}
