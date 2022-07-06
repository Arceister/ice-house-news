package service

import (
	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/usecase"
	"github.com/google/uuid"
)

type UsersService struct {
	usecase usecase.UsersUsecase
}

func NewUsersService(usecase usecase.UsersUsecase) UsersService {
	return UsersService{
		usecase: usecase,
	}
}

func (s UsersService) GetOneUserService(id string) (entity.User, error) {
	userData := entity.User{}
	var uid string

	err := s.usecase.GetOneUserUsecase(id).Scan(&uid,
		&userData.Email,
		&userData.Password,
		&userData.Name,
		&userData.Bio,
		&userData.Web,
		&userData.Picture,
	)
	if err != nil {
		return entity.User{}, err
	}
	userData.Id = uuid.Must(uuid.Parse(uid))

	return userData, err
}
