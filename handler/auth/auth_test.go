package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Arceister/ice-house-news/entity"
	serviceMock "github.com/Arceister/ice-house-news/service/mock"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandler_UserSignInHandler(t *testing.T) {
	mockUserService := new(serviceMock.UsersServiceMock)

	mockJSONRequest := `
	{
		"email": "testemail@email.com",
		"password": "123"
	}
	`
	mockAuthReturn := entity.UserAuthenticationReturn{
		Token:     "token",
		Scheme:    "bearer",
		ExpiresAt: time.Time{},
	}

	mockHandler := NewAuthHandler(mockUserService)

	t.Run("Success", func(t *testing.T) {
		type successStruct struct {
			Success bool                            `json:"success"`
			Message string                          `json:"message"`
			Data    entity.UserAuthenticationReturn `json:"data"`
		}

		userSignInRequest := entity.UserSignInRequest{
			Email:    "testemail@email.com",
			Password: "123",
		}

		req, err := http.NewRequest(http.MethodPost, "http://localhost:5055/api/auth/login", bytes.NewBufferString(mockJSONRequest))
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()

		mockUserService.On("SignInService", userSignInRequest).
			Return(mockAuthReturn, nil).Once()

		mockHandler.UserSignInHandler(w, req)

		var httpResponse successStruct
		jsonUnmarshalErr := json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
		if jsonUnmarshalErr != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, httpResponse)
		assert.Nil(t, err)
		assert.EqualValues(t, http.StatusOK, w.Code)
		assert.EqualValues(t, true, httpResponse.Success)
		assert.EqualValues(t, "Login successful", httpResponse.Message)

		mockUserService.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		type errorStruct struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}

		userSignInRequest := entity.UserSignInRequest{
			Email:    "testemail@email.com",
			Password: "123",
		}

		req, err := http.NewRequest(http.MethodPost, "http://localhost:5055/api/auth/login", bytes.NewBufferString(mockJSONRequest))
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()

		mockUserService.On("SignInService", userSignInRequest).
			Return(entity.UserAuthenticationReturn{}, errorUtils.NewUnauthorizedError("invalid")).Once()

		mockHandler.UserSignInHandler(w, req)

		var httpResponse errorStruct
		jsonUnmarshalErr := json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
		if jsonUnmarshalErr != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, httpResponse)
		assert.Nil(t, err)
		assert.EqualValues(t, http.StatusUnauthorized, w.Code)
		assert.EqualValues(t, false, httpResponse.Success)
		assert.EqualValues(t, "invalid", httpResponse.Message)

		mockUserService.AssertExpectations(t)
	})
}

func TestAuthHandler_ExtendToken(t *testing.T) {
	mockUserService := new(serviceMock.UsersServiceMock)

	mockAuthReturn := entity.UserAuthenticationReturn{
		Token:     "token",
		Scheme:    "bearer",
		ExpiresAt: time.Time{},
	}
	mockJWTClaims := jwt.MapClaims{
		"id": "10adc3ce-62e5-4b0a-82e0-fad9cc4b2c37",
	}
	userId := "10adc3ce-62e5-4b0a-82e0-fad9cc4b2c37"

	req, err := http.NewRequest(http.MethodGet, "http://localhost:5055/api/auth/token", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockHandler := NewAuthHandler(mockUserService)

	t.Run("Success", func(t *testing.T) {
		type successStruct struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}

		w := httptest.NewRecorder()
		req = req.WithContext(context.WithValue(context.Background(), "JWTProps", mockJWTClaims))

		mockUserService.On("ExtendToken", userId).
			Return(mockAuthReturn, nil).Once()

		mockHandler.ExtendTokenHandler(w, req)

		var httpResponse successStruct
		jsonUnmarshalErr := json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
		if jsonUnmarshalErr != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, httpResponse)
		assert.Nil(t, err)
		assert.EqualValues(t, http.StatusOK, w.Code)
		assert.EqualValues(t, true, httpResponse.Success)
		assert.EqualValues(t, "Login successful", httpResponse.Message)

		mockUserService.AssertExpectations(t)
	})
}
