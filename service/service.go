package service

import (
	"github.com/Arceister/ice-house-news/entity"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
)

type IUsersService interface {
	GetOneUserService(string) (entity.User, errorUtils.IErrorMessage)
	SignInService(entity.UserSignInRequest) (entity.UserAuthenticationReturn, errorUtils.IErrorMessage)
	CreateUserService(entity.User) errorUtils.IErrorMessage
	UpdateUserService(string, entity.User) errorUtils.IErrorMessage
	DeleteUserService(string) errorUtils.IErrorMessage
}

type ICategoriesService interface {
	GetAllNewsCategoryService() ([]entity.Categories, errorUtils.IErrorMessage)
	CreateCategoryService(string) errorUtils.IErrorMessage
}

type INewsService interface {
	GetNewsListService() ([]entity.NewsListOutput, errorUtils.IErrorMessage)
	GetNewsDetailService(string) (entity.NewsDetail, errorUtils.IErrorMessage)
	InsertNewsService(string, entity.NewsInputRequest) errorUtils.IErrorMessage
	UpdateNewsService(string, string, entity.NewsInputRequest) errorUtils.IErrorMessage
	DeleteNewsService(string, string) errorUtils.IErrorMessage
}

type ICommentService interface {
	GetCommentsOnNewsService(newsId string) ([]entity.Comment, errorUtils.IErrorMessage)
	InsertCommentService(commentRequest entity.CommentInsertRequest, newsId string, userId string) errorUtils.IErrorMessage
}
