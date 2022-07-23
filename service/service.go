package service

import "github.com/Arceister/ice-house-news/entity"

type IUsersService interface {
	GetOneUserService(string) (entity.User, error)
	SignInService(entity.UserSignInRequest) (*string, error)
	CreateUserService(entity.User) error
	UpdateUserService(string, entity.User) error
	DeleteUserService(string) error
}

type ICategoriesService interface {
	CreateCategoryService(string) error
}

type INewsService interface {
	GetNewsListService() ([]entity.NewsListOutput, error)
	GetNewsDetailService(string) (entity.NewsDetail, error)
	InsertNewsService(string, entity.NewsInputRequest) error
	UpdateNewsService(string, string, entity.NewsInputRequest) error
	DeleteNewsService(string, string) error
}
