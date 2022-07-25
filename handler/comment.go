package handler

import (
	"net/http"

	"github.com/Arceister/ice-house-news/server"
	"github.com/Arceister/ice-house-news/service"
	"github.com/go-chi/chi/v5"
)

type CommentHandler struct {
	service service.CommentService
}

func NewCommentHandler(service service.CommentService) CommentHandler {
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
