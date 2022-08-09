package service

import (
	"github.com/Arceister/ice-house-news/entity"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
)

var (
	GetOneUser func(string) (entity.User, errorUtils.IErrorMessage)
	SignIn     func(entity.UserSignInRequest) (entity.UserAuthenticationReturn, errorUtils.IErrorMessage)
	Extend     func(userID string) (entity.UserAuthenticationReturn, errorUtils.IErrorMessage)
	CreateUser func(entity.User) errorUtils.IErrorMessage
	UpdateUser func(string, entity.User) errorUtils.IErrorMessage
	DeleteUser func(string) errorUtils.IErrorMessage
)

type UsersServiceMock struct{}

func (m *UsersServiceMock) GetOneUserService(id string) (entity.User, errorUtils.IErrorMessage) {
	return GetOneUser(id)
}

func (m *UsersServiceMock) SignInService(userInput entity.UserSignInRequest) (entity.UserAuthenticationReturn, errorUtils.IErrorMessage) {
	return SignIn(userInput)
}

func (m *UsersServiceMock) ExtendToken(userID string) (entity.UserAuthenticationReturn, errorUtils.IErrorMessage) {
	return Extend(userID)
}

func (m *UsersServiceMock) CreateUserService(userData entity.User) errorUtils.IErrorMessage {
	return CreateUser(userData)
}

func (m *UsersServiceMock) UpdateUserService(id string, userData entity.User) errorUtils.IErrorMessage {
	return UpdateUser(id, userData)
}

func (m *UsersServiceMock) DeleteUserService(id string) errorUtils.IErrorMessage {
	return DeleteUser(id)
}
