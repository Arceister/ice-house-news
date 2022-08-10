package service

import (
	"testing"
	"time"

	"github.com/Arceister/ice-house-news/entity"
	repositoryMock "github.com/Arceister/ice-house-news/repository/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewsService_GetNewsList(t *testing.T) {
	mockNewsRepo := new(repositoryMock.NewsRepositoryMock)
	mockUsersRepo := new(repositoryMock.UserRepositoryMock)
	mockCategoriesRepo := new(repositoryMock.CategoriesRepositoryMock)

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

	mockNewsService := NewNewsService(mockNewsRepo, mockUsersRepo, mockCategoriesRepo)

	t.Run("Success", func(t *testing.T) {
		mockNewsRepo.On("GetNewsListRepository", 3, "Howak").
			Return(mockNewsList, nil).Once()

		news, err := mockNewsService.GetNewsListService("top_news", "Howak")
		if err != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, news)
		assert.Nil(t, err)

		mockNewsRepo.AssertExpectations(t)
	})
}
