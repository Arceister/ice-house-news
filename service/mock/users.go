package service

import (
	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/service"
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

type serviceMock struct{}

func NewServiceMock() service.IUsersService {
	return serviceMock{}
}

func (m serviceMock) GetOneUserService(id string) (entity.User, errorUtils.IErrorMessage) {
	return GetOneUser(id)
}

func (m serviceMock) SignInService(userInput entity.UserSignInRequest) (entity.UserAuthenticationReturn, errorUtils.IErrorMessage) {
	return SignIn(userInput)
}

func (m serviceMock) ExtendToken(userID string) (entity.UserAuthenticationReturn, errorUtils.IErrorMessage) {
	return Extend(userID)
}

func (m serviceMock) CreateUserService(userData entity.User) errorUtils.IErrorMessage {
	return CreateUser(userData)
}

func (m serviceMock) UpdateUserService(id string, userData entity.User) errorUtils.IErrorMessage {
	return UpdateUser(id, userData)
}

func (m serviceMock) DeleteUserService(id string) errorUtils.IErrorMessage {
	return DeleteUser(id)
}
