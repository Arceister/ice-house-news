package entity

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Bio      string    `json:"bio"`
	Web      string    `json:"web"`
	Picture  string    `json:"picture"`
}

type UserSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type IUsersService interface {
	GetOneUserService(id string) (User, error)
	CreateUserService(userData User) error
	UpdateUserService(id string, userData User) error
	DeleteUserService(id string) error
}

type IUsersUsecase interface {
	GetOneUserUsecase(id string) (User, error)
	CreateUserUsecase(id uuid.UUID, userData User) error
	UpdateUserUsecase(id string, userData User) error
	DeleteUserUsecase(id string) error
}
