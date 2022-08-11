package service

import (
	"github.com/Arceister/ice-house-news/entity"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
	"github.com/stretchr/testify/mock"
)

type CategoriesServiceMock struct {
	mock.Mock
}

func (m *CategoriesServiceMock) GetAllNewsCategoryService() ([]entity.Categories, errorUtils.IErrorMessage) {
	args := m.Called()
	if args.Get(1) == nil {
		return args.Get(0).([]entity.Categories), nil
	}
	return nil, args.Get(1).(errorUtils.IErrorMessage)
}

func (m *CategoriesServiceMock) CreateCategoryService(name string) errorUtils.IErrorMessage {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errorUtils.IErrorMessage)
}
