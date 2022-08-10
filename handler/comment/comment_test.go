package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Arceister/ice-house-news/entity"
	serviceMock "github.com/Arceister/ice-house-news/service/mock"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCommentHandler_GetComments(t *testing.T) {
	mockCommentService := new(serviceMock.CommentServiceMock)

	mockComment := entity.Comment{
		Id:          uuid.MustParse("8a950b11-8037-4ad6-81fc-8e53cb0c670d"),
		Description: "This is a comment.",
		User: entity.Commentator{
			Id:      uuid.MustParse("d73a3c2d-b34a-48dc-8b25-e9c164355bc8"),
			Name:    "Name",
			Picture: "String",
		},
		CreatedAt: time.Time{},
	}
	mockAllComment := []entity.Comment{mockComment}
	newsId := "8a950b11-8037-4ad6-81fc-8e53cb0c670d"

	mockHandler := NewCommentHandler(mockCommentService)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:5055/api/news/%s/comment", newsId), nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Success", func(t *testing.T) {
		type successStruct struct {
			Success bool             `json:"success"`
			Message string           `json:"message"`
			Data    []entity.Comment `json:"data"`
		}

		w := httptest.NewRecorder()
		routerCtx := chi.NewRouteContext()
		routerCtx.URLParams.Add("newsId", newsId)
		req = req.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, routerCtx))

		mockCommentService.On("GetCommentsOnNewsService", newsId).
			Return(mockAllComment, nil).Once()

		mockHandler.GetCommentsOnNewsHandler(w, req)

		var httpResponse successStruct
		err := json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
		if err != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, httpResponse)
		assert.Nil(t, err)
		assert.EqualValues(t, http.StatusOK, w.Code)
		assert.EqualValues(t, true, httpResponse.Success)
		assert.EqualValues(t, "comments get", httpResponse.Message)
		assert.EqualValues(t, mockAllComment, httpResponse.Data)

		mockCommentService.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		type errorStruct struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}

		w := httptest.NewRecorder()
		routerCtx := chi.NewRouteContext()
		routerCtx.URLParams.Add("newsId", newsId)
		req = req.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, routerCtx))

		mockCommentService.On("GetCommentsOnNewsService", newsId).
			Return(nil, errorUtils.NewInternalServerError("error message")).Once()

		mockHandler.GetCommentsOnNewsHandler(w, req)

		var httpResponse errorStruct
		err := json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
		if err != nil {
			t.Fatal(err)
		}

		assert.Nil(t, err)
		assert.EqualValues(t, http.StatusInternalServerError, w.Code)
		assert.EqualValues(t, false, httpResponse.Success)
		assert.EqualValues(t, "error message", httpResponse.Message)

		mockCommentService.AssertExpectations(t)
	})
}

func TestCommentHandler_InsertComment(t *testing.T) {
	mockCommentService := new(serviceMock.CommentServiceMock)

	mockJSONRequest := `{"description": "Insert comment"}`
	mockJWTClaims := jwt.MapClaims{
		"id": "10adc3ce-62e5-4b0a-82e0-fad9cc4b2c37",
	}
	mockComment := entity.CommentInsertRequest{
		Description: "Insert comment",
	}
	newsId := "8a950b11-8037-4ad6-81fc-8e53cb0c670d"
	userId := "10adc3ce-62e5-4b0a-82e0-fad9cc4b2c37"

	mockHandler := NewCommentHandler(mockCommentService)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:5055/api/news/%s/comment", newsId), bytes.NewBufferString(mockJSONRequest))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	routerCtx := chi.NewRouteContext()
	routerCtx.URLParams.Add("newsId", newsId)

	firstContext := context.WithValue(context.Background(), chi.RouteCtxKey, routerCtx)

	// req = req.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, routerCtx))
	req = req.WithContext(context.WithValue(firstContext, "JWTProps", mockJWTClaims))

	t.Run("Success", func(t *testing.T) {
		type successStruct struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}

		mockCommentService.On("InsertCommentService", mockComment, newsId, userId).
			Return(nil).Once()

		mockHandler.InsertCommentHandler(w, req)

		var httpResponse successStruct
		err := json.Unmarshal([]byte(w.Body.Bytes()), &httpResponse)
		if err != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, httpResponse)
		assert.Nil(t, err)
		assert.EqualValues(t, http.StatusOK, w.Code)
		assert.EqualValues(t, true, httpResponse.Success)
		assert.EqualValues(t, "insert comment success", httpResponse.Message)

		mockCommentService.AssertExpectations(t)
	})
}
