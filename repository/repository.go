package repository

import (
	"github.com/Arceister/ice-house-news/entity"
	"github.com/google/uuid"

	errorUtils "github.com/Arceister/ice-house-news/utils/error"
)

type IUsersRepository interface {
	GetOneUserRepository(string) (entity.User, errorUtils.IErrorMessage)
	GetUserByEmailRepository(string) (entity.User, errorUtils.IErrorMessage)
	CreateUserRepository(uuid.UUID, entity.User) errorUtils.IErrorMessage
	UpdateUserRepository(string, entity.User) errorUtils.IErrorMessage
	DeleteUserRepository(string) errorUtils.IErrorMessage
}

type ICategoriesRepository interface {
	GetAllNewsCategoryRepository() ([]entity.Categories, errorUtils.IErrorMessage)
	CreateCategoryRepository(categoryData entity.Categories) errorUtils.IErrorMessage
	CreateAndReturnCategoryRepository(entity.Categories) (uuid.UUID, errorUtils.IErrorMessage)
	GetCategoryByNameRepository(string) (entity.Categories, errorUtils.IErrorMessage)
}

type INewsRepository interface {
	GetNewsListRepository(int, string) ([]entity.NewsListOutput, errorUtils.IErrorMessage)
	GetNewsDetailRepository(string) (entity.NewsDetail, errorUtils.IErrorMessage)
	GetNewsUserRepository(string) (string, errorUtils.IErrorMessage)
	AddNewNewsRepository(entity.NewsInsert) errorUtils.IErrorMessage
	UpdateNewsRepository(entity.NewsInsert) errorUtils.IErrorMessage
	DeleteNewsRepository(string) errorUtils.IErrorMessage
}

type ICommentRepository interface {
	GetCommentsOnNewsRepository(string) ([]entity.Comment, errorUtils.IErrorMessage)
	InsertCommentRepository(entity.CommentInsert) errorUtils.IErrorMessage
}
