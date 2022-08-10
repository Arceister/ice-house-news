package service

import (
	"testing"

	"github.com/Arceister/ice-house-news/entity"
	repositoryMock "github.com/Arceister/ice-house-news/repository/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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
	})

	categories, err := mockCategoriesService.GetAllNewsCategoryService()
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, categories)
	assert.Nil(t, err)
	assert.EqualValues(t, mockCategories, categories)

	mockCategoriesRepo.AssertExpectations(t)
}
