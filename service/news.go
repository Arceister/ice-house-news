package service

import "github.com/Arceister/ice-house-news/repository"

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
