package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email" validate:"email, required"`
	Password string    `json:"password" validate:"required"`
	Name     string    `json:"name" validate:"required"`
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

type UserAuthenticationReturn struct {
	Token     string    `json:"token"`
	Scheme    string    `json:"scheme"`
	ExpiresAt time.Time `json:"expires_at"`
}
