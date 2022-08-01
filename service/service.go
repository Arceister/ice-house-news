package service

import "github.com/Arceister/ice-house-news/entity"

type IUsersService interface {
	GetOneUserService(string) (entity.User, error)
	SignInService(entity.UserSignInRequest) (entity.UserAuthenticationReturn, error)
	CreateUserService(entity.User) error
	UpdateUserService(string, entity.User) error
	DeleteUserService(string) error
}

type ICategoriesService interface {
	GetAllNewsCategoryService() ([]entity.Categories, error)
	CreateCategoryService(string) error
}

type INewsService interface {
	GetNewsListService() ([]entity.NewsListOutput, error)
	GetNewsDetailService(string) (entity.NewsDetail, error)
	InsertNewsService(string, entity.NewsInputRequest) error
	UpdateNewsService(string, string, entity.NewsInputRequest) error
	DeleteNewsService(string, string) error
}

type ICommentService interface {
	GetCommentsOnNewsService(newsId string) ([]entity.Comment, error)
	InsertCommentService(commentRequest entity.CommentInsertRequest, newsId string, userId string) error
}
