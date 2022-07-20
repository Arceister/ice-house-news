package service

import (
	"errors"
	"reflect"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/repository"
	"github.com/google/uuid"
)

type NewsService struct {
	newsRepository       repository.NewsRepository
	usersRepository      repository.UsersRepository
	categoriesRepository repository.CategoriesRepository
}

func NewNewsService(
	newsRepository repository.NewsRepository,
	usersRepository repository.UsersRepository,
	categoriesRepository repository.CategoriesRepository,
) NewsService {
	return NewsService{
		newsRepository:       newsRepository,
		usersRepository:      usersRepository,
		categoriesRepository: categoriesRepository,
	}
}

func (s NewsService) GetNewsListService() ([]entity.NewsListOutput, error) {
	return s.newsRepository.GetNewsListRepository()
}

func (s NewsService) GetNewsDetailService(newsId string) (entity.NewsDetail, error) {
	newsDetail, err := s.newsRepository.GetNewsDetailRepository(newsId)

	if reflect.DeepEqual(newsDetail, entity.NewsDetail{}) {
		return newsDetail, errors.New("news not found")
	}

	return newsDetail, err
}

func (s NewsService) InsertNewsService(userId string, newsInputData entity.NewsInputRequest) error {
	var newsInsertData entity.NewsInsert

	newsInsertData.NewsInputRequest = newsInputData

	newsUUID := uuid.Must(uuid.NewRandom())
	parsedUserUUID := uuid.Must(uuid.Parse(userId))

	categoryDetail, err := s.categoriesRepository.GetCategoryByNameRepository(newsInputData.Category)

	if categoryDetail == (entity.Categories{}) && err.Error() == "no rows in result set" {
		newCategoryUUID := uuid.Must(uuid.NewRandom())

		newCategoryData := entity.Categories{}
		newCategoryData.Id = newCategoryUUID
		newCategoryData.Name = &newsInputData.Category

		newCategoryId, err := s.categoriesRepository.CreateAndReturnCategoryRepository(newCategoryData)

		if err != nil {
			return err
		}

		categoryDetail.Id = *newCategoryId
	}

	if err != nil && err.Error() != "no rows in result set" {
		return err
	}

	newsInsertData.Id = newsUUID
	newsInsertData.UserId = parsedUserUUID
	newsInsertData.CategoryId = categoryDetail.Id

	err = s.newsRepository.AddNewNewsRepository(newsInsertData)
	if err != nil {
		return err
	}

	return err
}
