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
	args := m.Called(id)
	if args.Get(1) == nil {
		return args.Get(0).(entity.User), nil
	}
	return args.Get(0).(entity.User), args.Get(1).(errorUtils.IErrorMessage)
}

func (m *UserRepositoryMock) GetUserByEmailRepository(email string) (entity.User, errorUtils.IErrorMessage) {
	args := m.Called(email)
	if args.Get(1) == nil {
		return args.Get(0).(entity.User), nil
	}
	return args.Get(0).(entity.User), args.Get(1).(errorUtils.IErrorMessage)
}

func (m *UserRepositoryMock) CreateUserRepository(id uuid.UUID, userInput entity.User) errorUtils.IErrorMessage {
	args := m.Called(id, userInput)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errorUtils.IErrorMessage)
}

func (m *UserRepositoryMock) UpdateUserRepository(id string, userInput entity.User) errorUtils.IErrorMessage {
	args := m.Called(id, userInput)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errorUtils.IErrorMessage)
}

func (m *UserRepositoryMock) DeleteUserRepository(id string) errorUtils.IErrorMessage {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errorUtils.IErrorMessage)
}
