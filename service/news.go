package service

import (
	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/repository"
)

type NewsService struct {
	newsRepository  repository.NewsRepository
	usersRepository repository.UsersRepository
}

func NewNewsService(
	newsRepository repository.NewsRepository,
	usersRepository repository.UsersRepository,
) NewsService {
	return NewsService{
		newsRepository:  newsRepository,
		usersRepository: usersRepository,
	}
}

func (s NewsService) GetNewsListService() ([]entity.NewsListOutput, error) {
	return s.newsRepository.GetNewsListRepository()
}
