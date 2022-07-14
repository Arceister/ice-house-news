package usecase

import (
	"github.com/Arceister/ice-house-news/entity"
	"github.com/google/uuid"
)

type IUsersUsecase interface {
	GetOneUserUsecase(id string) (entity.User, error)
	CreateUserUsecase(id uuid.UUID, userData entity.User) error
	UpdateUserUsecase(id string, userData entity.User) error
	DeleteUserUsecase(id string) error
}
