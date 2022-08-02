package service

import (
	"database/sql"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/middleware"
	"github.com/Arceister/ice-house-news/repository"
	"github.com/Arceister/ice-house-news/service"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
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

func (s UsersService) GetOneUserService(id string) (entity.User, errorUtils.IErrorMessage) {
	userData, errorMessage := s.repository.GetOneUserRepository(id)

	if errorMessage != nil {
		return entity.User{}, errorMessage
	}

	return userData, nil
}

func (s UsersService) SignInService(userInput entity.UserSignInRequest) (entity.UserAuthenticationReturn, errorUtils.IErrorMessage) {
	var signInSchema entity.UserAuthenticationReturn

	validate := validator.New()
	err := validate.Struct(userInput)
	if err != nil {
		return signInSchema, errorUtils.NewUnprocessableEntityError("please input email/password")
	}

	userData, errorMessage := s.repository.GetUserByEmailRepository(userInput.Email)

	if errorMessage != nil && errorMessage.Message() == sql.ErrNoRows.Error() {
		return signInSchema, errorUtils.NewNotFoundError("username/password not found")
	} else if errorMessage != nil {
		return signInSchema, errorMessage
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(userInput.Password))

	if err != nil {
		return signInSchema, errorUtils.NewUnauthorizedError("wrong password")
	}

	token, expire, err := s.middleware.GenerateNewToken(userData)

	if err != nil {
		return signInSchema, errorUtils.NewInternalServerError(err.Error())
	}

	signInSchema.Token = *token
	signInSchema.Scheme = "Bearer"
	signInSchema.ExpiresAt = expire

	return signInSchema, nil
}

func (s UsersService) ExtendToken(userID string) (entity.UserAuthenticationReturn, errorUtils.IErrorMessage) {
	var signInSchema entity.UserAuthenticationReturn

	userData, errorMessage := s.repository.GetOneUserRepository(userID)

	if errorMessage != nil && errorMessage.Message() == sql.ErrNoRows.Error() {
		return signInSchema, errorUtils.NewNotFoundError("user not found")
	} else if errorMessage != nil {
		return signInSchema, errorMessage
	}

	token, expire, err := s.middleware.GenerateNewToken(userData)

	if err != nil {
		return signInSchema, errorUtils.NewInternalServerError(err.Error())
	}

	signInSchema.Token = *token
	signInSchema.Scheme = "Bearer"
	signInSchema.ExpiresAt = expire

	return signInSchema, nil
}

func (s UsersService) CreateUserService(userData entity.User) errorUtils.IErrorMessage {
	uniqueUserId := uuid.Must(uuid.NewRandom())

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	userData.Password = string(hashedPassword)

	errorMessage := s.repository.CreateUserRepository(uniqueUserId, userData)

	if errorMessage != nil {
		return errorMessage
	}

	return nil
}

func (s UsersService) UpdateUserService(id string, userData entity.User) errorUtils.IErrorMessage {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	userData.Password = string(hashedPassword)

	errorMessage := s.repository.UpdateUserRepository(id, userData)
	if errorMessage != nil {
		return errorMessage
	}

	return nil
}

func (s UsersService) DeleteUserService(id string) errorUtils.IErrorMessage {
	errorMessage := s.repository.DeleteUserRepository(id)
	if errorMessage != nil {
		return errorMessage
	}

	return nil
}
