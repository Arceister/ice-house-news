package service

import (
	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/repository"
	"github.com/Arceister/ice-house-news/service"
	"github.com/google/uuid"

	errorUtils "github.com/Arceister/ice-house-news/utils/error"
)

type CategoriesService struct {
	repository repository.ICategoriesRepository
}

func NewCategoriesService(repository repository.ICategoriesRepository) service.ICategoriesService {
	return CategoriesService{
		repository: repository,
	}
}

func (s CategoriesService) GetAllNewsCategoryService() ([]entity.Categories, errorUtils.IErrorMessage) {
	categories, err := s.repository.GetAllNewsCategoryRepository()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s CategoriesService) CreateCategoryService(categoryName string) errorUtils.IErrorMessage {
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
