package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/handler"
	"github.com/Arceister/ice-house-news/server"
	"github.com/Arceister/ice-house-news/service"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
)

type CommentHandler struct {
	service service.ICommentService
}

func NewCommentHandler(service service.ICommentService) handler.ICommentHandler {
	return CommentHandler{
		service: service,
	}
}

func (h CommentHandler) GetCommentsOnNewsHandler(w http.ResponseWriter, r *http.Request) {
	newsId := chi.URLParam(r, "newsId")
	newsComments, err := h.service.GetCommentsOnNewsService(newsId)

	if err != nil {
		server.ResponseJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	if newsComments == nil {
		server.ResponseJSON(w, http.StatusOK, true, "no comment in this news")
		return
	}

	server.ResponseJSONData(w, http.StatusOK, true, "comments get", newsComments)
}

func (h CommentHandler) InsertCommentHandler(w http.ResponseWriter, r *http.Request) {
	var commentInputRequest entity.CommentInsertRequest

	newsId := chi.URLParam(r, "newsId")
	userId := r.Context().Value("JWTProps").(jwt.MapClaims)["id"].(string)

	json.NewDecoder(r.Body).Decode(&commentInputRequest)

	err := h.service.InsertCommentService(commentInputRequest, newsId, userId)
	if err != nil && err.Error() == "comment insert failed" {
		server.ResponseJSON(w, http.StatusUnprocessableEntity, false, err.Error())
		return
	} else if err != nil {
		server.ResponseJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	server.ResponseJSON(w, http.StatusOK, true, "insert comment success")
}
