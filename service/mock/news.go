package service

import (
	"github.com/Arceister/ice-house-news/entity"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
	"github.com/stretchr/testify/mock"
)

type NewsServiceMock struct {
	mock.Mock
}

func (m *NewsServiceMock) GetNewsListService(scope string, category string) ([]entity.NewsListOutput, errorUtils.IErrorMessage) {
	args := m.Called(scope, category)
	if args.Get(1) == nil {
		return args.Get(0).([]entity.NewsListOutput), nil
	}
	return nil, args.Get(1).(errorUtils.IErrorMessage)
}

func (m *NewsServiceMock) GetNewsDetailService(newsId string) (entity.NewsDetail, errorUtils.IErrorMessage) {
	args := m.Called(newsId)
	if args.Get(1) == nil {
		return args.Get(0).(entity.NewsDetail), nil
	}
	return args.Get(0).(entity.NewsDetail), args.Get(1).(errorUtils.IErrorMessage)
}

func (m *NewsServiceMock) InsertNewsService(userId string, newsInput entity.NewsInputRequest) errorUtils.IErrorMessage {
	args := m.Called(userId, newsInput)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errorUtils.IErrorMessage)
}

func (m *NewsServiceMock) UpdateNewsService(userId string, newsId string, newsInput entity.NewsInputRequest) errorUtils.IErrorMessage {
	args := m.Called(userId, newsId, newsInput)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errorUtils.IErrorMessage)
}

func (m *NewsServiceMock) DeleteNewsService(userId string, newsId string) errorUtils.IErrorMessage {
	args := m.Called(userId, newsId)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errorUtils.IErrorMessage)
}
