package service

import (
	"github.com/Arceister/ice-house-news/entity"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
	"github.com/stretchr/testify/mock"
)

var (
	GetOneUser func(string) (entity.User, errorUtils.IErrorMessage)
	SignIn     func(entity.UserSignInRequest) (entity.UserAuthenticationReturn, errorUtils.IErrorMessage)
	Extend     func(userID string) (entity.UserAuthenticationReturn, errorUtils.IErrorMessage)
	CreateUser func(entity.User) errorUtils.IErrorMessage
	UpdateUser func(string, entity.User) errorUtils.IErrorMessage
	DeleteUser func(string) errorUtils.IErrorMessage
)

type UsersServiceMock struct {
	mock.Mock
}

func (m *UsersServiceMock) GetOneUserService(id string) (entity.User, errorUtils.IErrorMessage) {
	args := m.Called(id)
	if args.Get(1) == nil {
		return args.Get(0).(entity.User), nil
	}
	return args.Get(0).(entity.User), args.Get(1).(errorUtils.IErrorMessage)
}

func (m *UsersServiceMock) SignInService(userInput entity.UserSignInRequest) (entity.UserAuthenticationReturn, errorUtils.IErrorMessage) {
	args := m.Called(userInput)
	if args.Get(1) == nil {
		return args.Get(0).(entity.UserAuthenticationReturn), nil
	}
	return args.Get(0).(entity.UserAuthenticationReturn), args.Get(1).(errorUtils.IErrorMessage)
}

func (m *UsersServiceMock) ExtendToken(userID string) (entity.UserAuthenticationReturn, errorUtils.IErrorMessage) {
	args := m.Called(userID)
	if args.Get(1) == nil {
		return args.Get(0).(entity.UserAuthenticationReturn), nil
	}
	return args.Get(0).(entity.UserAuthenticationReturn), args.Get(1).(errorUtils.IErrorMessage)
}

func (m *UsersServiceMock) CreateUserService(userData entity.User) errorUtils.IErrorMessage {
	args := m.Called(userData)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errorUtils.IErrorMessage)
}

func (m *UsersServiceMock) UpdateUserService(id string, userData entity.User) errorUtils.IErrorMessage {
	args := m.Called(id, userData)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errorUtils.IErrorMessage)
}

func (m *UsersServiceMock) DeleteUserService(id string) errorUtils.IErrorMessage {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(errorUtils.IErrorMessage)
}
