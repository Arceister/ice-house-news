package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/Arceister/ice-house-news/middleware"
	"github.com/Arceister/ice-house-news/service"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	getOneUser  func(string) (entity.User, errorUtils.IErrorMessage)
	signIn      func(entity.UserSignInRequest) (entity.UserAuthenticationReturn, errorUtils.IErrorMessage)
	extendToken func(userID string) (entity.UserAuthenticationReturn, errorUtils.IErrorMessage)
	createUser  func(entity.User) errorUtils.IErrorMessage
	updateUser  func(string, entity.User) errorUtils.IErrorMessage
	deleteUser  func(string) errorUtils.IErrorMessage
)

type serviceMock struct{}

func NewServiceMock() service.IUsersService {
	return serviceMock{}
}

func (m serviceMock) GetOneUserService(id string) (entity.User, errorUtils.IErrorMessage) {
	return getOneUser(id)
}

func (m serviceMock) SignInService(userInput entity.UserSignInRequest) (entity.UserAuthenticationReturn, errorUtils.IErrorMessage) {
	return signIn(userInput)
}

func (m serviceMock) ExtendToken(userID string) (entity.UserAuthenticationReturn, errorUtils.IErrorMessage) {
	return extendToken(userID)
}

func (m serviceMock) CreateUserService(userData entity.User) errorUtils.IErrorMessage {
	return createUser(userData)
}

func (m serviceMock) UpdateUserService(id string, userData entity.User) errorUtils.IErrorMessage {
	return updateUser(id, userData)
}

func (m serviceMock) DeleteUserService(id string) errorUtils.IErrorMessage {
	return deleteUser(id)
}

func TestUserHandler_GetOneUser_Success(t *testing.T) {
	mockService := NewServiceMock()

	type ResponseDataStruct struct {
		entity.User
	}

	type successStruct struct {
		Success bool               `json:"success"`
		Message string             `json:"message"`
		Data    ResponseDataStruct `json:"data"`
	}

	getOneUser = func(s string) (entity.User, errorUtils.IErrorMessage) {
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

	userId := "8db82f7e-5736-4430-a62c-2e735177d895"

	mockHandler := NewUsersHandler(mockService)
	req, err := http.NewRequest("GET", "http://localhost:5055/api/users/"+userId, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mockHandler.GetOneUserHandler(w, req)

	var httpResponse successStruct
	err = json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, httpResponse)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, w.Code)
	assert.EqualValues(t, true, httpResponse.Success)
	assert.EqualValues(t, "get one user", httpResponse.Message)
	assert.EqualValues(t, uuid.MustParse(userId), httpResponse.Data.Id)
	assert.EqualValues(t, "testemail@email.com", httpResponse.Data.Email)
	assert.EqualValues(t, "123", httpResponse.Data.Password)
	assert.EqualValues(t, "Jagad", httpResponse.Data.Name)
	assert.EqualValues(t, func(val string) *string { return &val }("Bio"), httpResponse.Data.Bio)
	assert.EqualValues(t, func(val string) *string { return &val }("Web"), httpResponse.Data.Web)
	assert.EqualValues(t, func(val string) *string { return &val }("Picture"), httpResponse.Data.Picture)
}

func TestUserHandler_GetOneUser_Error(t *testing.T) {
	mockService := NewServiceMock()

	type errorStruct struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	getOneUser = func(s string) (entity.User, errorUtils.IErrorMessage) {
		return entity.User{}, errorUtils.NewInternalServerError("error message")
	}

	userId := "8db82f7e-5736-4430-a62c-2e735177d895"

	mockHandler := NewUsersHandler(mockService)
	req, err := http.NewRequest("GET", "http://localhost:5055/api/users/"+userId, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mockHandler.GetOneUserHandler(w, req)

	var httpResponse errorStruct
	err = json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, httpResponse)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, w.Code)
	assert.EqualValues(t, false, httpResponse.Success)
	assert.EqualValues(t, "error message", httpResponse.Message)
}

func TestUserHandler_GetOwnProfile_Success(t *testing.T) {
	mockService := NewServiceMock()

	type ResponseDataStruct struct {
		entity.User
	}

	type successStruct struct {
		Success bool               `json:"success"`
		Message string             `json:"message"`
		Data    ResponseDataStruct `json:"data"`
	}

	getOneUser = func(s string) (entity.User, errorUtils.IErrorMessage) {
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

	userId := "8db82f7e-5736-4430-a62c-2e735177d895"

	mockHandler := NewUsersHandler(mockService)
	middlewareMock := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	fakeTokenString, _, err := middlewareMock.GenerateNewToken(
		entity.User{
			Id:       uuid.MustParse("8db82f7e-5736-4430-a62c-2e735177d895"),
			Email:    "testemail@email.com",
			Password: "123",
			Name:     "Jagad",
			Bio:      func(val string) *string { return &val }("Bio"),
			Web:      func(val string) *string { return &val }("Web"),
			Picture:  func(val string) *string { return &val }("Picture"),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	fakeToken, err := jwt.Parse(*fakeTokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})
	if err != nil {
		t.Fatal(err)
	}

	fakeClaims, _ := fakeToken.Claims.(jwt.MapClaims)

	req, err := http.NewRequest("GET", "http://localhost:5055/api/users/"+userId, nil)
	req = req.WithContext(context.WithValue(req.Context(), "JWTProps", fakeClaims))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mockHandler.GetOwnProfile(w, req)

	var httpResponse successStruct
	err = json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, httpResponse)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, w.Code)
	assert.EqualValues(t, true, httpResponse.Success)
	assert.EqualValues(t, "get profile", httpResponse.Message)
	assert.EqualValues(t, uuid.MustParse(userId), httpResponse.Data.Id)
	assert.EqualValues(t, "testemail@email.com", httpResponse.Data.Email)
	assert.EqualValues(t, "123", httpResponse.Data.Password)
	assert.EqualValues(t, "Jagad", httpResponse.Data.Name)
	assert.EqualValues(t, func(val string) *string { return &val }("Bio"), httpResponse.Data.Bio)
	assert.EqualValues(t, func(val string) *string { return &val }("Web"), httpResponse.Data.Web)
	assert.EqualValues(t, func(val string) *string { return &val }("Picture"), httpResponse.Data.Picture)
}

func TestUserHandler_GetOwnProfile_Error(t *testing.T) {
	mockService := NewServiceMock()

	type errorStruct struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	getOneUser = func(s string) (entity.User, errorUtils.IErrorMessage) {
		return entity.User{}, errorUtils.NewInternalServerError("error message")
	}

	userId := "8db82f7e-5736-4430-a62c-2e735177d895"

	mockHandler := NewUsersHandler(mockService)
	middlewareMock := middleware.NewMiddlewareJWT(lib.App{Port: ":5000", SecretKey: "SECRET"})

	fakeTokenString, _, err := middlewareMock.GenerateNewToken(
		entity.User{
			Id:       uuid.MustParse("8db82f7e-5736-4430-a62c-2e735177d895"),
			Email:    "testemail@email.com",
			Password: "123",
			Name:     "Jagad",
			Bio:      func(val string) *string { return &val }("Bio"),
			Web:      func(val string) *string { return &val }("Web"),
			Picture:  func(val string) *string { return &val }("Picture"),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	fakeToken, err := jwt.Parse(*fakeTokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})
	if err != nil {
		t.Fatal(err)
	}

	fakeClaims, _ := fakeToken.Claims.(jwt.MapClaims)

	req, err := http.NewRequest("GET", "http://localhost:5055/api/users/"+userId, nil)
	req = req.WithContext(context.WithValue(req.Context(), "JWTProps", fakeClaims))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mockHandler.GetOwnProfile(w, req)

	var httpResponse errorStruct
	err = json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, httpResponse)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, w.Code)
	assert.EqualValues(t, false, httpResponse.Success)
	assert.EqualValues(t, "error message", httpResponse.Message)
}

func TestUserHandler_CreateUser_Success(t *testing.T) {
	mockService := NewServiceMock()

	type successStruct struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	createUser = func(u entity.User) errorUtils.IErrorMessage {
		return nil
	}

	mockHandler := NewUsersHandler(mockService)

	jsonRequest := `{
    "email": "e@a.com",
    "password": "Hash",
    "name": "Nama",
    "bio": "Ini Bio",
    "web": "Ini Web",
    "picture": "Ini pict"
	}`

	req, err := http.NewRequest("POST", "http://localhost:5055/api/users/", bytes.NewBufferString(jsonRequest))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mockHandler.CreateUserHandler(w, req)

	var httpResponse successStruct
	err = json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, httpResponse)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, w.Code)
	assert.EqualValues(t, true, httpResponse.Success)
	assert.EqualValues(t, "create user success", httpResponse.Message)
}

func TestUserHandler_CreateUser_Error(t *testing.T) {
	mockService := NewServiceMock()

	type errorStruct struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	createUser = func(u entity.User) errorUtils.IErrorMessage {
		return errorUtils.NewInternalServerError("error message")
	}

	mockHandler := NewUsersHandler(mockService)

	jsonRequest := `{
    ""
	}`

	req, err := http.NewRequest("POST", "http://localhost:5055/api/users/", bytes.NewBufferString(jsonRequest))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mockHandler.CreateUserHandler(w, req)

	var httpResponse errorStruct
	err = json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, httpResponse)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, w.Code)
	assert.EqualValues(t, false, httpResponse.Success)
	assert.EqualValues(t, "error message", httpResponse.Message)
}
