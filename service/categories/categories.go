package service

import (
	"github.com/Arceister/ice-house-news/entity"
	categories "github.com/Arceister/ice-house-news/repository/categories"
	"github.com/google/uuid"
)

type CategoriesService struct {
	repository categories.CategoriesRepository
}

func NewCategoriesService(repository categories.CategoriesRepository) CategoriesService {
	return CategoriesService{
		repository: repository,
	}
}

func (s CategoriesService) GetAllNewsCategoryService() ([]entity.Categories, error) {
	categories, err := s.repository.GetAllNewsCategoryRepository()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s CategoriesService) CreateCategoryService(categoryName string) error {
	var categoriesData entity.Categories

	newUuid := uuid.Must(uuid.NewRandom())

	categoriesData.Id = newUuid
	categoriesData.Name = categoryName

	err := s.repository.CreateCategoryRepository(categoriesData)
	if err != nil {
		return err
	}

	return nil
}
