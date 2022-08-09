package repository

import (
	"github.com/Arceister/ice-house-news/entity"
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

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetOneUserRepository(id string) (entity.User, errorUtils.IErrorMessage) {
	return GetOneUser(id)
}

func (m *UserRepositoryMock) GetUserByEmailRepository(id string) (entity.User, errorUtils.IErrorMessage) {
	return GetUserByEmail(id)
}

func (m *UserRepositoryMock) CreateUserRepository(id uuid.UUID, userInput entity.User) errorUtils.IErrorMessage {
	return CreateUser(id, userInput)
}

func (m *UserRepositoryMock) UpdateUserRepository(id string, userInput entity.User) errorUtils.IErrorMessage {
	return UpdateUser(id, userInput)
}

func (m *UserRepositoryMock) DeleteUserRepository(id string) errorUtils.IErrorMessage {
	return DeleteUser(id)
}
