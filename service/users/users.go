package service

import (
	"errors"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/middleware"
	"github.com/Arceister/ice-house-news/repository"
	"github.com/Arceister/ice-house-news/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	repository repository.IUsersRepository
	middleware middleware.MiddlewareJWT
}

func NewUsersService(
	repository repository.IUsersRepository,
	middleware middleware.MiddlewareJWT,
) service.IUsersService {
	return UsersService{
		repository: repository,
		middleware: middleware,
	}
}

func (s UsersService) GetOneUserService(id string) (entity.User, error) {
	userData, err := s.repository.GetOneUserRepository(id)

	if err != nil {
		return entity.User{}, err
	}

	return userData, nil
}

func (s UsersService) SignInService(userInput entity.UserSignInRequest) (entity.UserAuthenticationReturn, error) {
	var signInSchema entity.UserAuthenticationReturn

	validate := validator.New()
	err := validate.Struct(userInput)
	if err != nil {
		return signInSchema, errors.New("please input email/password")
	}

	userData, err := s.repository.GetUserByEmailRepository(userInput.Email)

	if err != nil && err.Error() == "sql: no rows in result set" {
		return signInSchema, errors.New("user not found")
	} else if err != nil {
		return signInSchema, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(userInput.Password))

	if err != nil {
		return signInSchema, errors.New("wrong password")
	}

	token, expire, err := s.middleware.GenerateNewToken(userData)

	if err != nil {
		return signInSchema, err
	}

	signInSchema.Token = *token
	signInSchema.Scheme = "Bearer"
	signInSchema.ExpiresAt = expire

	return signInSchema, nil
}

func (s UsersService) CreateUserService(userData entity.User) error {
	uniqueUserId := uuid.Must(uuid.NewRandom())

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)
	if err != nil {
		return err
	}

	userData.Password = string(hashedPassword)

	err = s.repository.CreateUserRepository(uniqueUserId, userData)

	if err != nil {
		return err
	}

	return nil
}

func (s UsersService) UpdateUserService(id string, userData entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)
	if err != nil {
		return err
	}

	userData.Password = string(hashedPassword)

	err = s.repository.UpdateUserRepository(id, userData)
	if err != nil {
		return err
	}

	return nil
}

func (s UsersService) DeleteUserService(id string) error {
	err := s.repository.DeleteUserRepository(id)
	if err != nil {
		return err
	}

	return nil
}
