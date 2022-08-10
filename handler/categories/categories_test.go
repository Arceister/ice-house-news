package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Arceister/ice-house-news/entity"
	serviceMock "github.com/Arceister/ice-house-news/service/mock"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCategoriesHandler_GetNewsCategories(t *testing.T) {
	mockCategoriesService := new(serviceMock.CategoriesServiceMock)

	mockCategory := entity.Categories{
		Id:   uuid.MustParse("00ac3c8f-62b1-4d80-8c67-b448423048e9"),
		Name: "Howak",
	}
	mockCategories := []entity.Categories{mockCategory}

	mockHandler := NewCategoriesHandler(mockCategoriesService)

	req, err := http.NewRequest(http.MethodGet, "http://localhost:5055/api/news/category", nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Success", func(t *testing.T) {
		type successStruct struct {
			Success bool                `json:"success"`
			Message string              `json:"message"`
			Data    []entity.Categories `json:"data"`
		}

		w := httptest.NewRecorder()

		mockCategoriesService.On("GetAllNewsCategoryService").
			Return(mockCategories, nil).Once()

		mockHandler.GetAllNewsCategoryHandler(w, req)

		var httpResponse successStruct
		err := json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
		if err != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, httpResponse)
		assert.Nil(t, err)
		assert.EqualValues(t, http.StatusOK, w.Code)
		assert.EqualValues(t, true, httpResponse.Success)
		assert.EqualValues(t, "get all categories", httpResponse.Message)
		assert.EqualValues(t, mockCategories, httpResponse.Data)

		mockCategoriesService.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		type errorStruct struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}

		w := httptest.NewRecorder()

		mockCategoriesService.On("GetAllNewsCategoryService").
			Return(nil, errorUtils.NewInternalServerError("error message")).Once()

		mockHandler.GetAllNewsCategoryHandler(w, req)

		var httpResponse errorStruct
		err := json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
		if err != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, httpResponse)
		assert.Nil(t, err)
		assert.EqualValues(t, http.StatusInternalServerError, w.Code)
		assert.EqualValues(t, false, httpResponse.Success)
		assert.EqualValues(t, "error message", httpResponse.Message)

		mockCategoriesService.AssertExpectations(t)
	})
}
