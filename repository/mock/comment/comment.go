package repository

import (
	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/repository"

	errorUtils "github.com/Arceister/ice-house-news/utils/error"
)

var (
	GetComments   func(string) ([]entity.Comment, errorUtils.IErrorMessage)
	InsertComment func(entity.CommentInsert) errorUtils.IErrorMessage
)

type repositoryMock struct{}

func NewCommentRepositoryMock() repository.ICommentRepository {
	return repositoryMock{}
}

func (m repositoryMock) GetCommentsOnNewsRepository(newsId string) ([]entity.Comment, errorUtils.IErrorMessage) {
	return GetComments(newsId)
}

func (m repositoryMock) InsertCommentRepository(commentData entity.CommentInsert) errorUtils.IErrorMessage {
	return InsertComment(commentData)
}
