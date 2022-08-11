package repository

import (
	"github.com/Arceister/ice-house-news/entity"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
	"github.com/stretchr/testify/mock"
)

type NewsRepositoryMock struct {
	mock.Mock
}

func (m *NewsRepositoryMock) GetNewsListRepository(scope int, category string) ([]entity.NewsListOutput, errorUtils.IErrorMessage) {
	args := m.Called(scope, category)
	if args.Get(1) == nil {
		return args.Get(0).([]entity.NewsListOutput), nil
	}
	return nil, args.Get(1).(errorUtils.IErrorMessage)
}

func (m *NewsRepositoryMock) GetNewsDetailRepository(newsId string) (entity.NewsDetail, errorUtils.IErrorMessage) {
	args := m.Called(newsId)
	if args.Get(1) == nil {
		return args.Get(0).(entity.NewsDetail), nil
	}
	return args.Get(0).(entity.NewsDetail), args.Get(1).(errorUtils.IErrorMessage)
}

func (m *NewsRepositoryMock) GetNewsUserRepository(newsId string) (string, errorUtils.IErrorMessage) {
	args := m.Called(newsId)
	if args.Get(1) == nil {
		return args.Get(0).(string), nil
	}
	return args.Get(0).(string), args.Get(1).(errorUtils.IErrorMessage)
}

func (m *NewsRepositoryMock) AddNewNewsRepository(newsData entity.NewsInsert) errorUtils.IErrorMessage {
	args := m.Called(newsData)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errorUtils.IErrorMessage)
}

func (m *NewsRepositoryMock) UpdateNewsRepository(newsData entity.NewsInsert) errorUtils.IErrorMessage {
	args := m.Called(newsData)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errorUtils.IErrorMessage)
}

func (m *NewsRepositoryMock) DeleteNewsRepository(newsId string) errorUtils.IErrorMessage {
	args := m.Called(newsId)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errorUtils.IErrorMessage)
}
