package repository

import (
	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	errorUtils "github.com/Arceister/ice-house-news/utils/error"
)

var (
	GetOneUser     func(id string) (entity.User, errorUtils.IErrorMessage)
	GetUserByEmail func(email string) (entity.User, errorUtils.IErrorMessage)
	CreateUser     func(id uuid.UUID, userInput entity.User) errorUtils.IErrorMessage
	UpdateUser     func(id string, userInput entity.User) errorUtils.IErrorMessage
	DeleteUser     func(id string) errorUtils.IErrorMessage
)

type repositoryMock struct {
	mock.Mock
}

func NewRepositoryMock() repository.IUsersRepository {
	return repositoryMock{}
}

func (m repositoryMock) GetOneUserRepository(id string) (entity.User, errorUtils.IErrorMessage) {
	return GetOneUser(id)
}

func (m repositoryMock) GetUserByEmailRepository(id string) (entity.User, errorUtils.IErrorMessage) {
	return GetUserByEmail(id)
}

func (m repositoryMock) CreateUserRepository(id uuid.UUID, userInput entity.User) errorUtils.IErrorMessage {
	return CreateUser(id, userInput)
}

func (m repositoryMock) UpdateUserRepository(id string, userInput entity.User) errorUtils.IErrorMessage {
	return UpdateUser(id, userInput)
}

func (m repositoryMock) DeleteUserRepository(id string) errorUtils.IErrorMessage {
	return DeleteUser(id)
}
