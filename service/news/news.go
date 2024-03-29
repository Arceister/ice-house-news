package service

import (
	"database/sql"
	"fmt"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/repository"
	"github.com/Arceister/ice-house-news/service"
	"github.com/google/uuid"

	errorUtils "github.com/Arceister/ice-house-news/utils/error"
)

var _ service.INewsService = (*NewsService)(nil)

type NewsService struct {
	newsRepository       repository.INewsRepository
	usersRepository      repository.IUsersRepository
	categoriesRepository repository.ICategoriesRepository
}

func NewNewsService(
	newsRepository repository.INewsRepository,
	usersRepository repository.IUsersRepository,
	categoriesRepository repository.ICategoriesRepository,
) *NewsService {
	return &NewsService{
		newsRepository:       newsRepository,
		usersRepository:      usersRepository,
		categoriesRepository: categoriesRepository,
	}
}

func (s *NewsService) GetNewsListService(scopeQuery, categoryQuery string) ([]entity.NewsListOutput, errorUtils.IErrorMessage) {
	var scopeNumber int

	switch scopeQuery {
	case "top_news":
		scopeNumber = 3
	default:
		scopeNumber = 0
	}

	newsList, err := s.newsRepository.GetNewsListRepository(scopeNumber, categoryQuery)
	if err != nil {
		return nil, err
	}

	return newsList, nil
}

func (s *NewsService) GetNewsDetailService(newsId string) (entity.NewsDetail, errorUtils.IErrorMessage) {
	newsDetail, err := s.newsRepository.GetNewsDetailRepository(newsId)

	if err != nil {
		return entity.NewsDetail{}, err
	}

	return newsDetail, nil
}

func (s *NewsService) InsertNewsService(userId string, newsInputData entity.NewsInputRequest) errorUtils.IErrorMessage {
	var newsInsertData entity.NewsInsert

	newsInsertData.NewsInputRequest = newsInputData

	newsUUID := uuid.Must(uuid.NewRandom())
	parsedUserUUID := uuid.Must(uuid.Parse(userId))

	categoryDetail, errorMessage := s.categoriesRepository.GetCategoryByNameRepository(newsInputData.Category)

	if errorMessage != nil && errorMessage.Message() == sql.ErrNoRows.Error() {
		newCategoryUUID := uuid.Must(uuid.NewRandom())

		newCategoryData := entity.Categories{}
		newCategoryData.Id = newCategoryUUID
		newCategoryData.Name = newsInputData.Category

		newCategoryId, err := s.categoriesRepository.CreateAndReturnCategoryRepository(newCategoryData)

		if err != nil {
			return err
		}

		categoryDetail.Id = newCategoryId
	}

	if errorMessage != nil && errorMessage.Message() != sql.ErrNoRows.Error() {
		return errorMessage
	}

	newsInsertData.Id = newsUUID
	newsInsertData.UserId = parsedUserUUID
	newsInsertData.CategoryId = categoryDetail.Id

	errorMessage = s.newsRepository.AddNewNewsRepository(newsInsertData)
	if errorMessage != nil {
		return errorMessage
	}

	return nil
}

func (s *NewsService) UpdateNewsService(
	userId string,
	newsId string,
	newsInputData entity.NewsInputRequest,
) errorUtils.IErrorMessage {
	var newsUpdateData entity.NewsInsert

	newsAuthorUUID, err := s.newsRepository.GetNewsUserRepository(newsId)
	if err != nil {
		return err
	}
	fmt.Println(newsAuthorUUID)
	fmt.Println(userId)

	if newsAuthorUUID != userId {
		return errorUtils.NewUnauthorizedError("user not authenticated")
	}

	categoryDetail, err := s.categoriesRepository.GetCategoryByNameRepository(newsInputData.Category)

	if err != nil && err.Message() == sql.ErrNoRows.Error() {
		newCategoryUUID := uuid.Must(uuid.NewRandom())

		newCategoryData := entity.Categories{}
		newCategoryData.Id = newCategoryUUID
		newCategoryData.Name = newsInputData.Category

		newCategoryId, err := s.categoriesRepository.CreateAndReturnCategoryRepository(newCategoryData)

		if err != nil {
			return err
		}

		categoryDetail.Id = newCategoryId
	}

	if err != nil && err.Message() != sql.ErrNoRows.Error() {
		return err
	}

	newsUpdateData.NewsInputRequest = newsInputData
	newsUpdateData.Id = uuid.Must(uuid.Parse(newsId))
	newsUpdateData.CategoryId = categoryDetail.Id

	err = s.newsRepository.UpdateNewsRepository(newsUpdateData)
	if err != nil {
		return err
	}

	return nil
}

func (s *NewsService) DeleteNewsService(
	userId string,
	newsId string,
) errorUtils.IErrorMessage {
	newsAuthorUUID, err := s.newsRepository.GetNewsUserRepository(newsId)
	if err != nil {
		return err
	}

	if newsAuthorUUID != userId {
		return errorUtils.NewUnauthorizedError("user not authenticated")
	}

	err = s.newsRepository.DeleteNewsRepository(newsId)
	if err != nil {
		return err
	}

	return nil
}
