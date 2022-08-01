package service

import (
	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/repository"
	"github.com/Arceister/ice-house-news/service"
	"github.com/google/uuid"
)

type CategoriesService struct {
	repository repository.ICategoriesRepository
}

func NewCategoriesService(repository repository.ICategoriesRepository) service.ICategoriesService {
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
