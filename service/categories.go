package service

import "github.com/Arceister/ice-house-news/repository"

type CategoriesService struct {
	repository repository.UsersRepository
}

func NewCategoriesService(repository repository.UsersRepository) CategoriesService {
	return CategoriesService{
		repository: repository,
	}
}
