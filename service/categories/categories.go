package service

import (
	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/service"
	"github.com/google/uuid"

	categoriesRepository "github.com/Arceister/ice-house-news/repository/categories"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
)

var _ service.ICategoriesService = (*CategoriesService)(nil)

type CategoriesService struct {
	repository categoriesRepository.CategoriesRepository
}

func NewCategoriesService(repository categoriesRepository.CategoriesRepository) CategoriesService {
	return CategoriesService{
		repository: repository,
	}
}

func (s *CategoriesService) GetAllNewsCategoryService() ([]entity.Categories, errorUtils.IErrorMessage) {
	categories, err := s.repository.GetAllNewsCategoryRepository()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoriesService) CreateCategoryService(categoryName string) errorUtils.IErrorMessage {
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
