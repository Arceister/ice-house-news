package repository

import (
	"github.com/Arceister/ice-house-news/entity"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type CategoriesRepositoryMock struct {
	mock.Mock
}

func (m *CategoriesRepositoryMock) GetAllNewsCategoryRepository() ([]entity.Categories, errorUtils.IErrorMessage) {
	args := m.Called()
	if args.Get(1) == nil {
		return args.Get(0).([]entity.Categories), nil
	}
	return nil, args.Get(1).(errorUtils.IErrorMessage)
}

func (m *CategoriesRepositoryMock) CreateCategoryRepository(categoryData entity.Categories) errorUtils.IErrorMessage {
	args := m.Called(categoryData)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errorUtils.IErrorMessage)
}

func (m *CategoriesRepositoryMock) CreateAndReturnCategoryRepository(categoryData entity.Categories) (uuid.UUID, errorUtils.IErrorMessage) {
	args := m.Called(categoryData)
	if args.Get(1) == nil {
		return args.Get(0).(uuid.UUID), nil
	}
	return uuid.Nil, args.Get(1).(errorUtils.IErrorMessage)
}

func (m *CategoriesRepositoryMock) GetCategoryByNameRepository(name string) (entity.Categories, errorUtils.IErrorMessage) {
	args := m.Called(name)
	if args.Get(1) == nil {
		return args.Get(0).(entity.Categories), nil
	}
	return args.Get(0).(entity.Categories), args.Get(1).(errorUtils.IErrorMessage)
}
