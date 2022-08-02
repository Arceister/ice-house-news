package service

import (
	"database/sql"
	"net/http"
	"testing"
	"time"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/Arceister/ice-house-news/middleware"
	"github.com/Arceister/ice-house-news/repository"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	currentTime    = time.Now()
	getOneUser     func(id string) (entity.User, errorUtils.IErrorMessage)
	getUserByEmail func(email string) (entity.User, errorUtils.IErrorMessage)
	createUser     func(id uuid.UUID, userInput entity.User) errorUtils.IErrorMessage
	updateUser     func(id string, userInput entity.User) errorUtils.IErrorMessage
	deleteUser     func(id string) errorUtils.IErrorMessage
)

type repositoryMock struct{}

func NewRepositoryMock() repository.IUsersRepository {
	return repositoryMock{}
}

func (m repositoryMock) GetOneUserRepository(id string) (entity.User, errorUtils.IErrorMessage) {
	return getOneUser(id)
}

func (m repositoryMock) GetUserByEmailRepository(id string) (entity.User, errorUtils.IErrorMessage) {
	return getUserByEmail(id)
}

func (m repositoryMock) CreateUserRepository(id uuid.UUID, userInput entity.User) errorUtils.IErrorMessage {
	return createUser(id, userInput)
}

func (m repositoryMock) UpdateUserRepository(id string, userInput entity.User) errorUtils.IErrorMessage {
	return updateUser(id, userInput)
}

func (m repositoryMock) DeleteUserRepository(id string) errorUtils.IErrorMessage {
	return deleteUser(id)
}

func TestUsersService_GetOneUser_Success(t *testing.T) {
	mockRepository := NewRepositoryMock()
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	getOneUser = func(id string) (entity.User, errorUtils.IErrorMessage) {
		return entity.User{
			Id:       uuid.MustParse("8db82f7e-5736-4430-a62c-2e735177d895"),
			Email:    "testemail@email.com",
			Password: "123",
			Name:     "Jagad",
			Bio:      func(val string) *string { return &val }("Bio"),
			Web:      func(val string) *string { return &val }("Web"),
			Picture:  func(val string) *string { return &val }("Picture"),
		}, nil
	}

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
	mockRepository := NewRepositoryMock()
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	getOneUser = func(id string) (entity.User, errorUtils.IErrorMessage) {
		return entity.User{}, errorUtils.NewNotFoundError("user not found")
	}

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
	mockRepository := NewRepositoryMock()
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	getUserByEmail = func(email string) (entity.User, errorUtils.IErrorMessage) {
		return entity.User{
			Id:       uuid.MustParse("8db82f7e-5736-4430-a62c-2e735177d895"),
			Email:    "testemail@email.com",
			Password: "$2a$10$nwPZff/XE2WrNKYzT6IHGepOaFFS1fJrP9jGXKWQ.JjX7YNlGPr6m",
			Name:     "Jagad",
			Bio:      func(val string) *string { return &val }("Bio"),
			Web:      func(val string) *string { return &val }("Web"),
			Picture:  func(val string) *string { return &val }("Picture"),
		}, nil
	}

	userSignIn := entity.UserSignInRequest{
		Email:    "testemail@email.com",
		Password: "123",
	}

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
	mockRepository := NewRepositoryMock()
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	getUserByEmail = func(email string) (entity.User, errorUtils.IErrorMessage) {
		return entity.User{
			Id:       uuid.MustParse("8db82f7e-5736-4430-a62c-2e735177d895"),
			Email:    "testemail@email.com",
			Password: "$2a$10$nwPZff/XE2WrNKYzT6IHGepOaFFS1fJrP9jGXKWQ.JjX7YNlGPr6m",
			Name:     "Jagad",
			Bio:      func(val string) *string { return &val }("Bio"),
			Web:      func(val string) *string { return &val }("Web"),
			Picture:  func(val string) *string { return &val }("Picture"),
		}, nil
	}

	userSignIn := entity.UserSignInRequest{
		Email:    "testemail@email.com",
		Password: "1234",
	}

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
	mockRepository := NewRepositoryMock()
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	getUserByEmail = func(email string) (entity.User, errorUtils.IErrorMessage) {
		return entity.User{
			Id:       uuid.MustParse("8db82f7e-5736-4430-a62c-2e735177d895"),
			Email:    "testemail@email.com",
			Password: "$2a$10$nwPZff/XE2WrNKYzT6IHGepOaFFS1fJrP9jGXKWQ.JjX7YNlGPr6m",
			Name:     "Jagad",
			Bio:      func(val string) *string { return &val }("Bio"),
			Web:      func(val string) *string { return &val }("Web"),
			Picture:  func(val string) *string { return &val }("Picture"),
		}, nil
	}

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
	mockRepository := NewRepositoryMock()
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	getUserByEmail = func(email string) (entity.User, errorUtils.IErrorMessage) {
		return entity.User{}, errorUtils.NewNotFoundError(sql.ErrNoRows.Error())
	}

	userSignIn := entity.UserSignInRequest{
		Email:    "testemailinvalid@email.com",
		Password: "123",
	}

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
	mockRepository := NewRepositoryMock()
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	getUserByEmail = func(email string) (entity.User, errorUtils.IErrorMessage) {
		return entity.User{}, errorUtils.NewInternalServerError("error message")
	}

	userSignIn := entity.UserSignInRequest{
		Email:    "testemailinvalid@email.com",
		Password: "123",
	}

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
	mockRepository := NewRepositoryMock()
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	getOneUser = func(id string) (entity.User, errorUtils.IErrorMessage) {
		return entity.User{
			Id:       uuid.MustParse("8db82f7e-5736-4430-a62c-2e735177d895"),
			Email:    "testemail@email.com",
			Password: "123",
			Name:     "Jagad",
			Bio:      func(val string) *string { return &val }("Bio"),
			Web:      func(val string) *string { return &val }("Web"),
			Picture:  func(val string) *string { return &val }("Picture"),
		}, nil
	}

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
	mockRepository := NewRepositoryMock()
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	getOneUser = func(id string) (entity.User, errorUtils.IErrorMessage) {
		return entity.User{}, errorUtils.NewNotFoundError(sql.ErrNoRows.Error())
	}

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
	mockRepository := NewRepositoryMock()
	middleware := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	getOneUser = func(id string) (entity.User, errorUtils.IErrorMessage) {
		return entity.User{}, errorUtils.NewInternalServerError("error message")
	}

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
