package service

import (
	"testing"

	"github.com/Arceister/ice-house-news/entity"
	repositoryMock "github.com/Arceister/ice-house-news/repository/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	errorUtils "github.com/Arceister/ice-house-news/utils/error"
)

func TestCategoriesService_GetAllNewsCategory(t *testing.T) {
	mockCategoriesRepo := new(repositoryMock.CategoriesRepositoryMock)

	mockCategory := entity.Categories{
		Id:   uuid.MustParse("00ac3c8f-62b1-4d80-8c67-b448423048e9"),
		Name: "Howak",
	}
	mockCategories := []entity.Categories{mockCategory}

	mockCategoriesService := NewCategoriesService(mockCategoriesRepo)

	t.Run("Success", func(t *testing.T) {
		mockCategoriesRepo.On("GetAllNewsCategoryRepository").
			Return(mockCategories, nil).Once()

		categories, err := mockCategoriesService.GetAllNewsCategoryService()
		if err != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, categories)
		assert.Nil(t, err)
		assert.EqualValues(t, mockCategories, categories)

		mockCategoriesRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockCategoriesRepo.On("GetAllNewsCategoryRepository").
			Return(nil, errorUtils.NewInternalServerError("error message")).Once()

		categories, err := mockCategoriesService.GetAllNewsCategoryService()
		if err == nil {
			t.Fatal("Error should be not null")
		}

		assert.Nil(t, categories)
		assert.NotNil(t, err)

		mockCategoriesRepo.AssertExpectations(t)
	})
}

func TestCategoriesService_CreateCategory(t *testing.T) {
	mockCategoriesRepo := new(repositoryMock.CategoriesRepositoryMock)
	mockCategoriesService := NewCategoriesService(mockCategoriesRepo)

	categoryName := "Howak"

	t.Run("Success", func(t *testing.T) {
		//Used mock.Anything because it generates random UUID
		mockCategoriesRepo.On("CreateCategoryRepository", mock.Anything).
			Return(nil).Once()

		err := mockCategoriesService.CreateCategoryService(categoryName)

		assert.Nil(t, err)

		mockCategoriesRepo.AssertExpectations(t)
	})

	t.Run("Success", func(t *testing.T) {
		//Used mock.Anything because it generates random UUID
		mockCategoriesRepo.On("CreateCategoryRepository", mock.Anything).
			Return(errorUtils.NewInternalServerError("error message")).Once()

		err := mockCategoriesService.CreateCategoryService(categoryName)

		assert.NotNil(t, err)

		mockCategoriesRepo.AssertExpectations(t)
	})
}
