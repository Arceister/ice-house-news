package service

import (
	"errors"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	repository repository.UsersRepository
}

func NewUsersService(repository repository.UsersRepository) UsersService {
	return UsersService{
		repository: repository,
	}
}

func (s UsersService) GetOneUserService(id string) (entity.User, error) {
	userData, err := s.repository.GetOneUserRepository(id)

	if userData == (entity.User{}) {
		return userData, errors.New("user not found")
	}

	return userData, err
}

func (s UsersService) CreateUserService(userData entity.User) error {
	uniqueUserId := uuid.Must(uuid.NewRandom())

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*userData.Password), 10)
	if err != nil {
		return err
	}

	*userData.Password = string(hashedPassword)

	return s.repository.CreateUserRepository(uniqueUserId, userData)
}

func (s UsersService) UpdateUserService(id string, userData entity.User) error {
	return s.repository.UpdateUserRepository(id, userData)
}

func (s UsersService) DeleteUserService(id string) error {
	return s.repository.DeleteUserRepository(id)
}
