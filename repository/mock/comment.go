package repository

import (
	"github.com/Arceister/ice-house-news/entity"
	"github.com/stretchr/testify/mock"

	errorUtils "github.com/Arceister/ice-house-news/utils/error"
)

type CommentRepositoryMock struct {
	mock.Mock
}

func (m *CommentRepositoryMock) GetCommentsOnNewsRepository(newsId string) ([]entity.Comment, errorUtils.IErrorMessage) {
	args := m.Called(newsId)
	if args.Get(1) == nil {
		return args.Get(0).([]entity.Comment), nil
	}
	return args.Get(0).([]entity.Comment), args.Get(1).(errorUtils.IErrorMessage)
}

func (m *CommentRepositoryMock) InsertCommentRepository(commentData entity.CommentInsert) errorUtils.IErrorMessage {
	args := m.Called(commentData)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errorUtils.IErrorMessage)
}
