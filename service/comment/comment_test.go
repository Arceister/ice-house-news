package service

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/Arceister/ice-house-news/entity"
	repositoryMock "github.com/Arceister/ice-house-news/repository/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestCommentService_GetComments(t *testing.T) {
	mockCommentRepo := new(repositoryMock.CommentRepositoryMock)
	mockNewsRepo := new(repositoryMock.NewsRepositoryMock)

	mockComment := entity.Comment{
		Id:          uuid.MustParse("8a950b11-8037-4ad6-81fc-8e53cb0c670d"),
		Description: "This is a comment.",
		User: entity.Commentator{
			Id:      uuid.MustParse("d73a3c2d-b34a-48dc-8b25-e9c164355bc8"),
			Name:    "Name",
			Picture: "String",
		},
		CreatedAt: time.Now(),
	}
	mockAllComment := []entity.Comment{mockComment}

	newsId := "8a950b11-8037-4ad6-81fc-8e53cb0c670d"

	mockService := NewCommentService(mockNewsRepo, mockCommentRepo)
	t.Run("Success", func(t *testing.T) {
		mockCommentRepo.On("GetCommentsOnNewsRepository", newsId).
			Return(mockAllComment, nil).Once()

		comment, err := mockService.GetCommentsOnNewsService(newsId)
		if err != nil {
			t.Fatal(err)
		}

		assert.Nil(t, err)
		assert.Equal(t, mockAllComment, comment)
	})
}
