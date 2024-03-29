package service

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/Arceister/ice-house-news/entity"
	repositoryMock "github.com/Arceister/ice-house-news/repository/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	errorUtils "github.com/Arceister/ice-house-news/utils/error"
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

		mockCommentRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockCommentRepo.On("GetCommentsOnNewsRepository", newsId).
			Return(nil, errorUtils.NewInternalServerError("error message")).Once()

		comment, err := mockService.GetCommentsOnNewsService(newsId)
		if err == nil {
			t.Fatal("Error should be not null")
		}

		assert.NotNil(t, err)
		assert.Nil(t, comment)

		mockCommentRepo.AssertExpectations(t)
	})
}

func TestCommentService_InsertComment(t *testing.T) {
	mockCommentRepo := new(repositoryMock.CommentRepositoryMock)
	mockNewsRepo := new(repositoryMock.NewsRepositoryMock)

	commentInsertMock := entity.CommentInsertRequest{
		Description: "Some comment mock.",
	}
	commentFailedMock := entity.CommentInsertRequest{
		Description: "",
	}
	newsId := "8a950b11-8037-4ad6-81fc-8e53cb0c670d"
	userId := "d73a3c2d-b34a-48dc-8b25-e9c164355bc8"

	mockService := NewCommentService(mockNewsRepo, mockCommentRepo)

	t.Run("Success", func(t *testing.T) {
		//Used mock.Anything because there's some random generation on UUID
		mockCommentRepo.On("InsertCommentRepository", mock.Anything).
			Return(nil).Once()

		err := mockService.InsertCommentService(commentInsertMock, newsId, userId)

		assert.Nil(t, err)

		mockCommentRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		//Used mock.Anything because there's some random generation on UUID
		mockCommentRepo.On("InsertCommentRepository", mock.Anything).
			Return(errorUtils.NewInternalServerError("error message")).Once()

		err := mockService.InsertCommentService(commentInsertMock, newsId, userId)

		assert.NotNil(t, err)

		mockCommentRepo.AssertExpectations(t)
	})

	t.Run("Validation Error", func(t *testing.T) {
		err := mockService.InsertCommentService(commentFailedMock, newsId, userId)

		assert.NotNil(t, err)

		mockCommentRepo.AssertExpectations(t)
	})
}
