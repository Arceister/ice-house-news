package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Arceister/ice-house-news/entity"
	serviceMock "github.com/Arceister/ice-house-news/service/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewsHandler_GetNewsList(t *testing.T) {
	mockNewsService := new(serviceMock.NewsServiceMock)

	mockNewsResult := entity.NewsListOutput{
		Id:               uuid.MustParse("922c7afd-643e-4e44-ab51-c80dc137674a"),
		Title:            "News Title",
		SlugUrl:          "news-title",
		CoverImage:       func(val string) *string { return &val }("Cover"),
		AdditionalImages: []string{"ABC"},
		CreatedAt:        time.Time{},
		Nsfw:             false,
		Category: entity.NewsCategory{
			Id:   uuid.MustParse("d414197c-0fa0-46c1-ac29-69c4cdc0ed11"),
			Name: "Howak",
		},
		Author: entity.NewsAuthor{
			Id:      uuid.MustParse("e65d7793-bcc6-467c-88b1-9636ee745f45"),
			Name:    "Name",
			Picture: func(val string) *string { return &val }("Picture"),
		},
		Counter: entity.NewsCounter{
			Upvote:   10,
			Downvote: 23,
			Comment:  2,
			View:     1000,
		},
	}
	mockNewsList := []entity.NewsListOutput{mockNewsResult}

	mockHandler := NewNewsHandler(mockNewsService)

	req, err := http.NewRequest(http.MethodGet, "http://localhost:5055/api/news", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	t.Run("Success", func(t *testing.T) {
		type successStruct struct {
			Success bool                    `json:"success"`
			Message string                  `json:"message"`
			Data    []entity.NewsListOutput `json:"data"`
		}

		mockNewsService.On("GetNewsListService", "", "").
			Return(mockNewsList, nil).Once()

		mockHandler.GetNewsListHandler(w, req)

		var httpResponse successStruct
		err := json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
		if err != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, httpResponse)
		assert.Nil(t, err)
		assert.EqualValues(t, http.StatusOK, w.Code)
		assert.EqualValues(t, true, httpResponse.Success)
		assert.EqualValues(t, "get news list", httpResponse.Message)
		assert.EqualValues(t, mockNewsList, httpResponse.Data)

		mockNewsService.AssertExpectations(t)
	})
}
