package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Arceister/ice-house-news/entity"
	serviceMock "github.com/Arceister/ice-house-news/service/mock"
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

	req, err := http.NewRequest(http.MethodPost, "http://localhost:5055/api/auth/login", bytes.NewBufferString(mockJSONRequest))
	if err != nil {
		t.Fatal(err)
	}

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
}
