package service

import "github.com/Arceister/ice-house-news/entity"

type IUsersService interface {
	GetOneUserService(id string) (entity.User, error)
	CreateUserService(userData entity.User) error
	UpdateUserService(id string, userData entity.User) error
	DeleteUserService(id string) error
}
