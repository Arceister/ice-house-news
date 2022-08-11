package service

import (
	"github.com/Arceister/ice-house-news/entity"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
	"github.com/stretchr/testify/mock"
)

type CommentServiceMock struct {
	mock.Mock
}

func (m *CommentServiceMock) GetCommentsOnNewsService(newsId string) ([]entity.Comment, errorUtils.IErrorMessage) {
	args := m.Called(newsId)
	if args.Get(1) == nil {
		return args.Get(0).([]entity.Comment), nil
	}
	return nil, args.Get(1).(errorUtils.IErrorMessage)
}

func (m *CommentServiceMock) InsertCommentService(commentRequest entity.CommentInsertRequest, newsId string, userId string) errorUtils.IErrorMessage {
	args := m.Called(commentRequest, newsId, userId)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errorUtils.IErrorMessage)
}
