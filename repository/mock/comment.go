package repository

import (
	"github.com/Arceister/ice-house-news/entity"

	errorUtils "github.com/Arceister/ice-house-news/utils/error"
)

var (
	GetComments   func(string) ([]entity.Comment, errorUtils.IErrorMessage)
	InsertComment func(entity.CommentInsert) errorUtils.IErrorMessage
)

type MockCommentRepository struct{}

func (m *MockCommentRepository) GetCommentsOnNewsRepository(newsId string) ([]entity.Comment, errorUtils.IErrorMessage) {
	return GetComments(newsId)
}

func (m *MockCommentRepository) InsertCommentRepository(commentData entity.CommentInsert) errorUtils.IErrorMessage {
	return InsertComment(commentData)
}
