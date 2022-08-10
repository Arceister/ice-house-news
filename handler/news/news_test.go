package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Arceister/ice-house-news/entity"
	serviceMock "github.com/Arceister/ice-house-news/service/mock"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
	"github.com/go-chi/chi/v5"
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

	t.Run("Success", func(t *testing.T) {
		type successStruct struct {
			Success bool                    `json:"success"`
			Message string                  `json:"message"`
			Data    []entity.NewsListOutput `json:"data"`
		}

		w := httptest.NewRecorder()

		q := req.URL.Query()
		q.Add("scope", "top_news")
		q.Add("category", "Howak")
		req.URL.RawQuery = q.Encode()

		mockNewsService.On("GetNewsListService", "top_news", "Howak").
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

	t.Run("Error", func(t *testing.T) {
		type errorStruct struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}

		w := httptest.NewRecorder()

		q := req.URL.Query()
		q.Add("scope", "top_news")
		q.Add("category", "Howak")
		req.URL.RawQuery = q.Encode()

		mockNewsService.On("GetNewsListService", "top_news", "Howak").
			Return(nil, errorUtils.NewInternalServerError("error message")).Once()

		mockHandler.GetNewsListHandler(w, req)

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

		mockNewsService.AssertExpectations(t)
	})
}

func TestNewsHandler_GetNewsDetail(t *testing.T) {
	mockNewsService := new(serviceMock.NewsServiceMock)

	mockNewsResult := entity.NewsDetail{
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
		Content: "Lorem Ipsum",
	}

	mockHandler := NewNewsHandler(mockNewsService)
	newsId := "922c7afd-643e-4e44-ab51-c80dc137674a"

	req, err := http.NewRequest(http.MethodGet, "http://localhost:5055/api/news"+newsId, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Success", func(t *testing.T) {
		type successStruct struct {
			Success bool              `json:"success"`
			Message string            `json:"message"`
			Data    entity.NewsDetail `json:"data"`
		}

		w := httptest.NewRecorder()
		routerCtx := chi.NewRouteContext()
		routerCtx.URLParams.Add("newsId", newsId)
		req = req.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, routerCtx))

		mockNewsService.On("GetNewsDetailService", newsId).
			Return(mockNewsResult, nil).Once()

		mockHandler.GetNewsDetailHandler(w, req)

		var httpResponse successStruct
		err := json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
		if err != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, httpResponse)
		assert.Nil(t, err)
		assert.EqualValues(t, http.StatusOK, w.Code)
		assert.EqualValues(t, true, httpResponse.Success)
		assert.EqualValues(t, "get news detail", httpResponse.Message)
		assert.EqualValues(t, mockNewsResult, httpResponse.Data)

		mockNewsService.AssertExpectations(t)
	})
}
