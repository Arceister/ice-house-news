package service

import (
	"errors"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/usecase"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"golang.org/x/crypto/bcrypt"
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
	userData, err := s.usecase.GetOneUserUsecase(id)

	if userData == (entity.User{}) {
		return userData, errors.New("user not found")
	}

	return userData, err
}

func (s UsersService) CreateUserService(userData entity.User) error {
	uniqueUserId := uuid.Must(uuid.NewRandom())

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)
	if err != nil {
		return err
	}

	userData.Password = string(hashedPassword)

	return s.usecase.CreateUserUsecase(uniqueUserId, userData)
}

func (s UsersService) UpdateUserService(uuid string, userData entity.User) error {
	return s.usecase.UpdateUserUsecase(uuid, userData)
}

func (s UsersService) DeleteUserService(uuid string) (pgconn.CommandTag, error) {
	return s.usecase.DeleteUserUsecase(uuid)
}
