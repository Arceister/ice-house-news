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

	if err != nil && err.Error() == "no rows in result set" {
		server.ResponseJSON(w, http.StatusNotFound, false, "news not found")
		return
	}
	if err != nil {
		server.ResponseJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	server.ResponseJSONData(w, http.StatusOK, true, "get news detail", newsDetail)
}

func (h NewsHandler) AddNewNewsHandler(w http.ResponseWriter, r *http.Request) {
	var newsInput entity.NewsInputRequest

	currentUserId := r.Context().Value("JWTProps").(jwt.MapClaims)["id"].(string)
	json.NewDecoder(r.Body).Decode(&newsInput)

	err := h.service.InsertNewsService(currentUserId, newsInput)

	if err != nil && (err.Error() == "news not created" ||
		err.Error() == "input additional image failed" ||
		err.Error() == "input news counter failed") {
		server.ResponseJSON(w, http.StatusUnprocessableEntity, false, err.Error())
		return
	}
	if err != nil {
		server.ResponseJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	server.ResponseJSON(w, http.StatusOK, true, "insert new news success")
}

func (h NewsHandler) UpdateNewsHandler(w http.ResponseWriter, r *http.Request) {
	var newsInput entity.NewsInputRequest

	currentUserId := r.Context().Value("JWTProps").(jwt.MapClaims)["id"].(string)
	newsId := chi.URLParam(r, "newsId")

	json.NewDecoder(r.Body).Decode(&newsInput)

	err := h.service.UpdateNewsService(currentUserId, newsId, newsInput)

	if err != nil && err.Error() == "news not updated" {
		server.ResponseJSON(w, http.StatusUnprocessableEntity, false, err.Error())
		return
	}
	if err != nil && err.Error() == "no rows in result set" {
		server.ResponseJSON(w, http.StatusNotFound, false, "news not found")
		return
	}
	if err != nil && err.Error() == "user not authenticated" {
		server.ResponseJSON(w, http.StatusUnauthorized, false, err.Error())
		return
	}
	if err != nil {
		server.ResponseJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	server.ResponseJSON(w, http.StatusOK, true, "news updated")
}

func (h NewsHandler) DeleteNewsHandler(w http.ResponseWriter, r *http.Request) {
	currentUserId := r.Context().Value("JWTProps").(jwt.MapClaims)["id"].(string)
	newsId := chi.URLParam(r, "newsId")

	err := h.service.DeleteNewsService(currentUserId, newsId)

	if err != nil && err.Error() == "news not deleted" {
		server.ResponseJSON(w, http.StatusUnprocessableEntity, false, err.Error())
		return
	}
	if err != nil && err.Error() == "no rows in result set" {
		server.ResponseJSON(w, http.StatusNotFound, false, "news not found")
		return
	}
	if err != nil && err.Error() == "user not authenticated" {
		server.ResponseJSON(w, http.StatusUnauthorized, false, err.Error())
		return
	}
	if err != nil {
		server.ResponseJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	server.ResponseJSON(w, http.StatusOK, true, "news deleted")
}
