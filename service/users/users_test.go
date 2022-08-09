package service

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/Arceister/ice-house-news/middleware"
	userRepositoryMock "github.com/Arceister/ice-house-news/repository/mock"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUsersService_GetOneUser_Success(t *testing.T) {
	mockRepository := new(userRepositoryMock.UserRepositoryMock)

	mockResult := entity.User{
		Id:       uuid.MustParse("8db82f7e-5736-4430-a62c-2e735177d895"),
		Email:    "testemail@email.com",
		Password: "123",
		Name:     "Jagad",
		Bio:      func(val string) *string { return &val }("Bio"),
		Web:      func(val string) *string { return &val }("Web"),
		Picture:  func(val string) *string { return &val }("Picture"),
	}

	mockRepository.On("GetOneUserRepository", "8db82f7e-5736-4430-a62c-2e735177d895").
		Return(mockResult, nil)
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	mockService := NewUsersService(mockRepository, middleware)
	userData, err := mockService.GetOneUserService("8db82f7e-5736-4430-a62c-2e735177d895")
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, userData)
	assert.Nil(t, err)
	assert.EqualValues(t, userData.Id, uuid.MustParse("8db82f7e-5736-4430-a62c-2e735177d895"))
	assert.EqualValues(t, userData.Email, "testemail@email.com")
	assert.EqualValues(t, userData.Name, "Jagad")
	assert.EqualValues(t, userData.Bio, func(val string) *string { return &val }("Bio"))
	assert.EqualValues(t, userData.Web, func(val string) *string { return &val }("Web"))
	assert.EqualValues(t, userData.Picture, func(val string) *string { return &val }("Picture"))
}

func TestUsersService_GetOneUser_Failed(t *testing.T) {
	mockRepository := new(userRepositoryMock.UserRepositoryMock)
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	mockRepository.On("GetOneUserRepository", "8db82f7e-5736-4430-a62c-2e735177d895").
		Return(entity.User{}, errorUtils.NewNotFoundError("user not found"))

	mockService := NewUsersService(mockRepository, middleware)
	userData, err := mockService.GetOneUserService("8db82f7e-5736-4430-a62c-2e735177d895")
	if err == nil {
		t.Fatal("Didn't get errors")
	}

	assert.NotNil(t, err)
	assert.Equal(t, userData, entity.User{})
	assert.EqualValues(t, err.Message(), "user not found")
	assert.EqualValues(t, err.Status(), http.StatusNotFound)
}

func TestUsersService_SignIn_Success(t *testing.T) {
	mockRepository := new(userRepositoryMock.UserRepositoryMock)
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	mockResult := entity.User{
		Id:       uuid.MustParse("8db82f7e-5736-4430-a62c-2e735177d895"),
		Email:    "testemail@email.com",
		Password: "$2a$10$nwPZff/XE2WrNKYzT6IHGepOaFFS1fJrP9jGXKWQ.JjX7YNlGPr6m",
		Name:     "Jagad",
		Bio:      func(val string) *string { return &val }("Bio"),
		Web:      func(val string) *string { return &val }("Web"),
		Picture:  func(val string) *string { return &val }("Picture"),
	}

	userSignIn := entity.UserSignInRequest{
		Email:    "testemail@email.com",
		Password: "123",
	}

	mockRepository.On("GetUserByEmailRepository", "testemail@email.com").
		Return(mockResult, nil).Once()

	mockService := NewUsersService(mockRepository, middleware)
	userAuthenticationReturn, err := mockService.SignInService(userSignIn)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, userAuthenticationReturn)
	assert.Nil(t, err)
	assert.EqualValues(t, userAuthenticationReturn.Scheme, "Bearer")
	assert.NotNil(t, userAuthenticationReturn)
}

func TestUsersService_SignIn_Unauthorized(t *testing.T) {
	mockRepository := new(userRepositoryMock.UserRepositoryMock)
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	mockResult := entity.User{
		Id:       uuid.MustParse("8db82f7e-5736-4430-a62c-2e735177d895"),
		Email:    "testemail@email.com",
		Password: "$2a$10$nwPZff/XE2WrNKYzT6IHGepOaFFS1fJrP9jGXKWQ.JjX7YNlGPr6m",
		Name:     "Jagad",
		Bio:      func(val string) *string { return &val }("Bio"),
		Web:      func(val string) *string { return &val }("Web"),
		Picture:  func(val string) *string { return &val }("Picture"),
	}

	userSignIn := entity.UserSignInRequest{
		Email:    "testemail@email.com",
		Password: "1234",
	}

	mockRepository.On("GetUserByEmailRepository", "testemail@email.com").
		Return(mockResult, nil).Once()

	mockService := NewUsersService(mockRepository, middleware)
	userAuthenticationReturn, err := mockService.SignInService(userSignIn)
	if err == nil {
		t.Fatal("No error detected, fail.")
	}

	assert.NotNil(t, err)
	assert.EqualValues(t, userAuthenticationReturn, entity.UserAuthenticationReturn{})
	assert.EqualValues(t, err.Status(), http.StatusUnauthorized)
	assert.EqualValues(t, err.Message(), "wrong password")
}

func TestUsersService_SignIn_UnprocessableEntity(t *testing.T) {
	mockRepository := new(userRepositoryMock.UserRepositoryMock)
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	userSignIn := entity.UserSignInRequest{
		Email:    "",
		Password: "",
	}

	mockService := NewUsersService(mockRepository, middleware)
	userAuthenticationReturn, err := mockService.SignInService(userSignIn)
	if err == nil {
		t.Fatal("No error detected, fail.")
	}

	assert.NotNil(t, err)
	assert.EqualValues(t, userAuthenticationReturn, entity.UserAuthenticationReturn{})
	assert.EqualValues(t, err.Status(), http.StatusUnprocessableEntity)
	assert.EqualValues(t, err.Message(), "please input email/password")
}

func TestUsersService_SignIn_NotFound(t *testing.T) {
	mockRepository := new(userRepositoryMock.UserRepositoryMock)
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	userSignIn := entity.UserSignInRequest{
		Email:    "testemailinvalid@email.com",
		Password: "123",
	}

	mockRepository.On("GetUserByEmailRepository", "testemailinvalid@email.com").
		Return(entity.User{}, errorUtils.NewNotFoundError(sql.ErrNoRows.Error())).Once()

	mockService := NewUsersService(mockRepository, middleware)
	userAuthenticationReturn, err := mockService.SignInService(userSignIn)
	if err == nil {
		t.Fatal("No error detected, fail.")
	}

	assert.NotNil(t, err)
	assert.EqualValues(t, userAuthenticationReturn, entity.UserAuthenticationReturn{})
	assert.EqualValues(t, err.Status(), http.StatusNotFound)
	assert.EqualValues(t, err.Message(), "username/password not found")
}

func TestUsersService_SignIn_InternalServerError(t *testing.T) {
	mockRepository := new(userRepositoryMock.UserRepositoryMock)
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	userSignIn := entity.UserSignInRequest{
		Email:    "testemail@email.com",
		Password: "123",
	}

	mockRepository.On("GetUserByEmailRepository", "testemail@email.com").
		Return(entity.User{}, errorUtils.NewInternalServerError("error message")).Once()

	mockService := NewUsersService(mockRepository, middleware)
	userAuthenticationReturn, err := mockService.SignInService(userSignIn)
	if err == nil {
		t.Fatal("No error detected, fail.")
	}

	assert.NotNil(t, err)
	assert.EqualValues(t, userAuthenticationReturn, entity.UserAuthenticationReturn{})
	assert.EqualValues(t, err.Status(), http.StatusInternalServerError)
	assert.EqualValues(t, err.Message(), "error message")
}

func TestUsersService_ExtendToken_Success(t *testing.T) {
	mockRepository := new(userRepositoryMock.UserRepositoryMock)
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	mockResult := entity.User{
		Id:       uuid.MustParse("8db82f7e-5736-4430-a62c-2e735177d895"),
		Email:    "testemail@email.com",
		Password: "123",
		Name:     "Jagad",
		Bio:      func(val string) *string { return &val }("Bio"),
		Web:      func(val string) *string { return &val }("Web"),
		Picture:  func(val string) *string { return &val }("Picture"),
	}

	mockRepository.On("GetOneUserRepository", "8db82f7e-5736-4430-a62c-2e735177d895").
		Return(mockResult, nil)

	mockService := NewUsersService(mockRepository, middleware)
	userAuthenticationReturn, err := mockService.ExtendToken("8db82f7e-5736-4430-a62c-2e735177d895")
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, userAuthenticationReturn)
	assert.Nil(t, err)
	assert.EqualValues(t, userAuthenticationReturn.Scheme, "Bearer")
	assert.NotNil(t, userAuthenticationReturn)
}

func TestUsersService_ExtendToken_NotFound(t *testing.T) {
	mockRepository := new(userRepositoryMock.UserRepositoryMock)
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	mockRepository.On("GetOneUserRepository", "8db82f7e-5736-4430-a62c-2e735177d895").
		Return(entity.User{}, errorUtils.NewNotFoundError(sql.ErrNoRows.Error()))

	mockService := NewUsersService(mockRepository, middleware)
	userAuthenticationReturn, err := mockService.ExtendToken("8db82f7e-5736-4430-a62c-2e735177d895")
	if err == nil {
		t.Fatal(err)
	}

	assert.NotNil(t, err)
	assert.EqualValues(t, userAuthenticationReturn, entity.UserAuthenticationReturn{})
	assert.EqualValues(t, err.Status(), http.StatusNotFound)
	assert.EqualValues(t, err.Message(), "user not found")
}

func TestUsersService_ExtendToken_InternalServerError(t *testing.T) {
	mockRepository := new(userRepositoryMock.UserRepositoryMock)
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	mockRepository.On("GetOneUserRepository", "8db82f7e-5736-4430-a62c-2e735177d895").
		Return(entity.User{}, errorUtils.NewInternalServerError("error message"))

	mockService := NewUsersService(mockRepository, middleware)
	userAuthenticationReturn, err := mockService.ExtendToken("8db82f7e-5736-4430-a62c-2e735177d895")
	if err == nil {
		t.Fatal(err)
	}

	assert.NotNil(t, err)
	assert.EqualValues(t, userAuthenticationReturn, entity.UserAuthenticationReturn{})
	assert.EqualValues(t, err.Status(), http.StatusInternalServerError)
	assert.EqualValues(t, err.Message(), "error message")
}

func TestUsersService_CreateUser_Success(t *testing.T) {
	mockRepository := new(userRepositoryMock.UserRepositoryMock)
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	userData := entity.User{
		Email:    "testemail@email.com",
		Password: "123",
		Name:     "Jagad",
		Bio:      func(val string) *string { return &val }("Bio"),
		Web:      func(val string) *string { return &val }("Web"),
		Picture:  func(val string) *string { return &val }("Picture"),
	}

	//Used Mock.Anything because there are some hashing function.
	mockRepository.On("CreateUserRepository", mock.Anything, mock.Anything).
		Return(nil).Once()

	mockService := NewUsersService(mockRepository, middleware)
	err := mockService.CreateUserService(userData)

	assert.Nil(t, err)
}

func TestUsersService_CreateUser_Failed(t *testing.T) {
	mockRepository := new(userRepositoryMock.UserRepositoryMock)
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	userData := entity.User{
		Email:    "testemail@email.com",
		Password: "123",
		Name:     "Jagad",
		Bio:      func(val string) *string { return &val }("Bio"),
		Web:      func(val string) *string { return &val }("Web"),
		Picture:  func(val string) *string { return &val }("Picture"),
	}

	//Used Mock.Anything because there are some hashing function.
	mockRepository.On("CreateUserRepository", mock.Anything, mock.Anything).
		Return(errorUtils.NewInternalServerError("error message")).Once()

	mockService := NewUsersService(mockRepository, middleware)
	err := mockService.CreateUserService(userData)

	assert.NotNil(t, err)
	assert.EqualValues(t, err.Status(), http.StatusInternalServerError)
	assert.EqualValues(t, err.Message(), "error message")
}

func TestUsersService_UpdateUser_Success(t *testing.T) {
	mockRepository := new(userRepositoryMock.UserRepositoryMock)
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	userData := entity.User{
		Email:    "testemail@email.com",
		Password: "123",
		Name:     "Jagad",
		Bio:      func(val string) *string { return &val }("Bio"),
		Web:      func(val string) *string { return &val }("Web"),
		Picture:  func(val string) *string { return &val }("Picture"),
	}

	mockRepository.On("UpdateUserRepository", "8db82f7e-5736-4430-a62c-2e735177d895", mock.Anything).
		Return(nil).Once()

	mockService := NewUsersService(mockRepository, middleware)
	err := mockService.UpdateUserService("8db82f7e-5736-4430-a62c-2e735177d895", userData)

	assert.Nil(t, err)
}

func TestUsersService_UpdateUser_Failed(t *testing.T) {
	mockRepository := new(userRepositoryMock.UserRepositoryMock)
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	userData := entity.User{
		Id:       uuid.MustParse("8db82f7e-5736-4430-a62c-2e735177d895"),
		Email:    "testemail@email.com",
		Password: "123",
		Name:     "Jagad",
		Bio:      func(val string) *string { return &val }("Bio"),
		Web:      func(val string) *string { return &val }("Web"),
		Picture:  func(val string) *string { return &val }("Picture"),
	}

	mockRepository.On("UpdateUserRepository", "8db82f7e-5736-4430-a62c-2e735177d895", mock.Anything).
		Return(errorUtils.NewInternalServerError("error message")).Once()

	mockService := NewUsersService(mockRepository, middleware)
	err := mockService.UpdateUserService("8db82f7e-5736-4430-a62c-2e735177d895", userData)

	assert.NotNil(t, err)
	assert.EqualValues(t, err.Status(), http.StatusInternalServerError)
	assert.EqualValues(t, err.Message(), "error message")
}

func TestUsersService_DeleteUser_Success(t *testing.T) {
	mockRepository := new(userRepositoryMock.UserRepositoryMock)
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	mockRepository.On("DeleteUserRepository", "8db82f7e-5736-4430-a62c-2e735177d895").
		Return(nil).Once()

	mockService := NewUsersService(mockRepository, middleware)
	err := mockService.DeleteUserService("8db82f7e-5736-4430-a62c-2e735177d895")

	assert.Nil(t, err)
}

func TestUsersService_DeleteUser_Failed(t *testing.T) {
	mockRepository := new(userRepositoryMock.UserRepositoryMock)
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	mockRepository.On("DeleteUserRepository", "8db82f7e-5736-4430-a62c-2e735177d895").
		Return(errorUtils.NewInternalServerError("error message")).Once()

	mockService := NewUsersService(mockRepository, middleware)
	err := mockService.DeleteUserService("8db82f7e-5736-4430-a62c-2e735177d895")

	assert.NotNil(t, err)
	assert.EqualValues(t, err.Status(), http.StatusInternalServerError)
	assert.EqualValues(t, err.Message(), "error message")
}
