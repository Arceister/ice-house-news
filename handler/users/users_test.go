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
	usersServiceMock "github.com/Arceister/ice-house-news/service/mock"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_GetOneUser_Success(t *testing.T) {
	mockService := new(usersServiceMock.UsersServiceMock)

	type ResponseDataStruct struct {
		entity.User
	}

	type successStruct struct {
		Success bool               `json:"success"`
		Message string             `json:"message"`
		Data    ResponseDataStruct `json:"data"`
	}

	mockResult := entity.User{
		Id:       uuid.MustParse("8db82f7e-5736-4430-a62c-2e735177d895"),
		Email:    "testemail@email.com",
		Password: "123",
		Name:     "Jagad",
		Bio:      func(val string) *string { return &val }("Bio"),
		Web:      func(val string) *string { return &val }("Web"),
		Picture:  func(val string) *string { return &val }("Picture"),
	}

	userId := "8db82f7e-5736-4430-a62c-2e735177d895"

	mockHandler := NewUsersHandler(mockService)
	req, err := http.NewRequest(http.MethodGet, "http://localhost:5055/api/users/"+userId, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	routerCtx := chi.NewRouteContext()
	routerCtx.URLParams.Add("uuid", userId)
	req = req.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, routerCtx))

	mockService.On("GetOneUserService", userId).
		Return(mockResult, nil).Once()
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
	mockService := new(usersServiceMock.UsersServiceMock)

	type errorStruct struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	userId := "8db82f7e-5736-4430-a62c-2e735177d895"

	mockHandler := NewUsersHandler(mockService)
	req, err := http.NewRequest(http.MethodGet, "http://localhost:5055/api/users/"+userId, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	routerCtx := chi.NewRouteContext()
	routerCtx.URLParams.Add("uuid", userId)
	req = req.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, routerCtx))

	mockService.On("GetOneUserService", userId).
		Return(entity.User{}, errorUtils.NewInternalServerError("error message")).Once()
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
	mockService := new(usersServiceMock.UsersServiceMock)

	type ResponseDataStruct struct {
		entity.User
	}

	type successStruct struct {
		Success bool               `json:"success"`
		Message string             `json:"message"`
		Data    ResponseDataStruct `json:"data"`
	}

	mockResult := entity.User{
		Id:       uuid.MustParse("8db82f7e-5736-4430-a62c-2e735177d895"),
		Email:    "testemail@email.com",
		Password: "123",
		Name:     "Jagad",
		Bio:      func(val string) *string { return &val }("Bio"),
		Web:      func(val string) *string { return &val }("Web"),
		Picture:  func(val string) *string { return &val }("Picture"),
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

	req, err := http.NewRequest(http.MethodGet, "http://localhost:5055/api/users/", nil)
	req = req.WithContext(context.WithValue(req.Context(), "JWTProps", fakeClaims))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mockService.On("GetOneUserService", userId).
		Return(mockResult, nil).Once()

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
	mockService := new(usersServiceMock.UsersServiceMock)

	type errorStruct struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
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

	req, err := http.NewRequest(http.MethodGet, "http://localhost:5055/api/users/"+userId, nil)
	req = req.WithContext(context.WithValue(req.Context(), "JWTProps", fakeClaims))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mockService.On("GetOneUserService", userId).
		Return(entity.User{}, errorUtils.NewInternalServerError("error message")).Once()

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
	mockService := new(usersServiceMock.UsersServiceMock)

	type successStruct struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
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

	mockInput := entity.User{
		Email:    "e@a.com",
		Password: "Hash",
		Name:     "Nama",
		Bio:      func(val string) *string { return &val }("Ini Bio"),
		Web:      func(val string) *string { return &val }("Ini Web"),
		Picture:  func(val string) *string { return &val }("Ini pict"),
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:5055/api/users/", bytes.NewBufferString(jsonRequest))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mockService.On("CreateUserService", mockInput).
		Return(nil).Once()

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
	mockService := new(usersServiceMock.UsersServiceMock)

	type errorStruct struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
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

	mockInput := entity.User{
		Email:    "e@a.com",
		Password: "Hash",
		Name:     "Nama",
		Bio:      func(val string) *string { return &val }("Ini Bio"),
		Web:      func(val string) *string { return &val }("Ini Web"),
		Picture:  func(val string) *string { return &val }("Ini pict"),
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:5055/api/users/", bytes.NewBufferString(jsonRequest))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mockService.On("CreateUserService", mockInput).
		Return(errorUtils.NewInternalServerError("error message")).Once()

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

func TestUserHandler_UpdateUser_Success(t *testing.T) {
	mockService := new(usersServiceMock.UsersServiceMock)

	type successStruct struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	mockHandler := NewUsersHandler(mockService)

	userId := "8db82f7e-5736-4430-a62c-2e735177d895"

	jsonRequest := `{
    "email": "e@a.com",
    "password": "Hash",
    "name": "Nama",
    "bio": "Ini Bio",
    "web": "Ini Web",
    "picture": "Ini pict"
	}`

	mockInput := entity.User{
		Email:    "e@a.com",
		Password: "Hash",
		Name:     "Nama",
		Bio:      func(val string) *string { return &val }("Ini Bio"),
		Web:      func(val string) *string { return &val }("Ini Web"),
		Picture:  func(val string) *string { return &val }("Ini pict"),
	}

	req, err := http.NewRequest(http.MethodPut, "http://localhost:5055/api/users/"+userId, bytes.NewBufferString(jsonRequest))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	routerCtx := chi.NewRouteContext()
	routerCtx.URLParams.Add("uuid", userId)
	req = req.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, routerCtx))

	mockService.On("UpdateUserService", userId, mockInput).
		Return(nil).Once()

	mockHandler.UpdateUserHandler(w, req)

	var httpResponse successStruct
	err = json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, httpResponse)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, w.Code)
	assert.EqualValues(t, true, httpResponse.Success)
	assert.EqualValues(t, "update user success", httpResponse.Message)
}

func TestUserHandler_UpdateUser_Error(t *testing.T) {
	mockService := new(usersServiceMock.UsersServiceMock)

	type errorStruct struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	mockHandler := NewUsersHandler(mockService)

	userId := "8db82f7e-5736-4430-a62c-2e735177d895"

	jsonRequest := `{
    "email": "e@a.com",
    "password": "Hash",
    "name": "Nama",
    "bio": "Ini Bio",
    "web": "Ini Web",
    "picture": "Ini pict"
	}`

	mockInput := entity.User{
		Email:    "e@a.com",
		Password: "Hash",
		Name:     "Nama",
		Bio:      func(val string) *string { return &val }("Ini Bio"),
		Web:      func(val string) *string { return &val }("Ini Web"),
		Picture:  func(val string) *string { return &val }("Ini pict"),
	}

	req, err := http.NewRequest(http.MethodPut, "http://localhost:5055/api/users/"+userId, bytes.NewBufferString(jsonRequest))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	routerCtx := chi.NewRouteContext()
	routerCtx.URLParams.Add("uuid", userId)
	req = req.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, routerCtx))

	mockService.On("UpdateUserService", userId, mockInput).
		Return(errorUtils.NewInternalServerError("error message")).Once()

	mockHandler.UpdateUserHandler(w, req)

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

func TestUserHandler_DeleteUser_Success(t *testing.T) {
	mockService := new(usersServiceMock.UsersServiceMock)

	type successStruct struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	mockHandler := NewUsersHandler(mockService)

	userId := "8db82f7e-5736-4430-a62c-2e735177d895"

	req, err := http.NewRequest(http.MethodDelete, "http://localhost:5055/api/users/"+userId, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	routerCtx := chi.NewRouteContext()
	routerCtx.URLParams.Add("uuid", userId)
	req = req.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, routerCtx))

	mockService.On("DeleteUserService", userId).
		Return(nil).Once()

	mockHandler.DeleteUserHandler(w, req)

	var httpResponse successStruct
	err = json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, httpResponse)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, w.Code)
	assert.EqualValues(t, true, httpResponse.Success)
	assert.EqualValues(t, "delete user success", httpResponse.Message)
}

func TestUserHandler_DeleteUser_Error(t *testing.T) {
	mockService := new(usersServiceMock.UsersServiceMock)

	type errorStruct struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	mockHandler := NewUsersHandler(mockService)

	userId := "8db82f7e-5736-4430-a62c-2e735177d895"

	req, err := http.NewRequest(http.MethodDelete, "http://localhost:5055/api/users/"+userId, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	routerCtx := chi.NewRouteContext()
	routerCtx.URLParams.Add("uuid", userId)
	req = req.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, routerCtx))

	mockService.On("DeleteUserService", userId).
		Return(errorUtils.NewInternalServerError("error message")).Once()

	mockHandler.DeleteUserHandler(w, req)

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
