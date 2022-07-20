package repository

import (
	"github.com/Arceister/ice-house-news/entity"
	"github.com/google/uuid"
)

type IUsersRepository interface {
	GetOneUserRepository(string) (entity.User, error)
	CreateUserRepository(uuid.UUID, entity.User) error
	UpdateUserRepository(string, entity.User) error
	DeleteUserRepository(string) error
}

type ICategoriesRepository interface {
	CreateCategoryRepository(entity.Categories) error
	CreateAndReturnCategoryRepository(entity.Categories) (*uuid.UUID, error)
	GetCategoryByNameRepository(string) (entity.Categories, error)
}

type INewsRepository interface {
	GetNewsListRepository() ([]entity.NewsListOutput, error)
	GetNewsDetailRepository(string) (entity.NewsDetail, error)
	GetNewsUserRepository(string) (*string, error)
	AddNewNewsRepository(entity.NewsInsert) error
	UpdateNewNewsRepository(entity.NewsInsert) error
	DeleteNewsRepository(string) error
}
