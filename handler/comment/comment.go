package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/handler"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"

	response "github.com/Arceister/ice-house-news/server/response"
	commentService "github.com/Arceister/ice-house-news/service/comment"
)

var _ handler.ICommentHandler = (*CommentHandler)(nil)

type CommentHandler struct {
	service commentService.CommentService
}

func NewCommentHandler(service commentService.CommentService) CommentHandler {
	return CommentHandler{
		service: service,
	}
}

func (h CommentHandler) GetCommentsOnNewsHandler(w http.ResponseWriter, r *http.Request) {
	newsId := chi.URLParam(r, "newsId")
	newsComments, err := h.service.GetCommentsOnNewsService(newsId)

	if err != nil {
		response.ErrorResponse(w, err)
		return
	}

	response.SuccessResponseWithData(w, http.StatusOK, "comments get", newsComments)
}

func (h CommentHandler) InsertCommentHandler(w http.ResponseWriter, r *http.Request) {
	var commentInputRequest entity.CommentInsertRequest

	newsId := chi.URLParam(r, "newsId")
	userId := r.Context().Value("JWTProps").(jwt.MapClaims)["id"].(string)

	json.NewDecoder(r.Body).Decode(&commentInputRequest)

	err := h.service.InsertCommentService(commentInputRequest, newsId, userId)
	if err != nil {
		response.ErrorResponse(w, err)
		return
	}

	response.SuccessResponse(w, http.StatusOK, "insert comment success")
}
