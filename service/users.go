package service

import "github.com/Arceister/ice-house-news/usecase"

type UsersService struct {
	usecase usecase.UsersUsecase
}

func NewUsersService(usecase usecase.UsersUsecase) UsersService {
	return UsersService{
		usecase: usecase,
	}
}
