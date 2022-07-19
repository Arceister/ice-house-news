package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/server"
	"github.com/Arceister/ice-house-news/service"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
)

type NewsHandler struct {
	service service.NewsService
}

func NewNewsHandler(service service.NewsService) NewsHandler {
	return NewsHandler{
		service: service,
	}
}

func (h NewsHandler) GetNewsListHandler(w http.ResponseWriter, r *http.Request) {
	newsList, err := h.service.GetNewsListService()

	if err != nil {
		server.ResponseJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	server.ResponseJSONData(w, http.StatusOK, true, "get news list", newsList)
}

func (h NewsHandler) GetNewsDetailHandler(w http.ResponseWriter, r *http.Request) {
	newsId := chi.URLParam(r, "newsId")
	newsDetail, err := h.service.GetNewsDetailService(newsId)

	if err != nil && err.Error() == "news not found" {
		server.ResponseJSON(w, http.StatusNotFound, false, err.Error())
		return
	} else if err != nil {
		server.ResponseJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	server.ResponseJSONData(w, http.StatusOK, true, "get news detail", newsDetail)
}

func (h NewsHandler) AddNewNewsHandler(w http.ResponseWriter, r *http.Request) {
	var newsInput entity.NewsInsert

	currentUserId := r.Context().Value("JWTProps").(jwt.MapClaims)["id"].(string)
	json.NewDecoder(r.Body).Decode(&newsInput)

	err := h.service.InsertNewsService(currentUserId, newsInput)

	if err != nil {
		server.ResponseJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	server.ResponseJSON(w, http.StatusOK, true, "insert new news success")
}
