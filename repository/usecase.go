package repository

import (
	"github.com/Arceister/ice-house-news/entity"
	"github.com/google/uuid"
)

type IUsersRepository interface {
	GetOneUserRepository(id string) (entity.User, error)
	CreateUserRepository(id uuid.UUID, userData entity.User) error
	UpdateUserRepository(id string, userData entity.User) error
	DeleteUserRepository(id string) error
}
