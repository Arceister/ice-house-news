package service

import (
	"errors"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/middleware"
	"github.com/Arceister/ice-house-news/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	repository repository.UsersRepository
	middleware middleware.MiddlewareJWT
}

func NewUsersService(
	repository repository.UsersRepository,
	middleware middleware.MiddlewareJWT,
) UsersService {
	return UsersService{
		repository: repository,
		middleware: middleware,
	}
}

func (s UsersService) GetOneUserService(id string) (entity.User, error) {
	userData, err := s.repository.GetOneUserRepository(id)

	if userData == (entity.User{}) {
		return userData, errors.New("user not found")
	}

	return userData, err
}

func (s UsersService) SignInService(userInput entity.UserSignIn) (*string, error) {
	if userInput.Email == nil || userInput.Password == nil || len(*userInput.Email) == 0 || len(*userInput.Password) == 0 {
		return nil, errors.New("Email/Password Input Empty")
	}

	userData, err := s.repository.GetUserByEmailRepository(*userInput.Email)

	if err != nil && err.Error() == "no rows in result set" {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(*userData.Password), []byte(*userInput.Password))

	if err != nil {
		return nil, errors.New("wrong password")
	}

	token, err := s.middleware.GenerateNewToken(userData)

	if err != nil {
		return nil, err
	}

	return token, nil
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
