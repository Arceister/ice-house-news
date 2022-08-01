package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/handler"
	"github.com/Arceister/ice-house-news/service"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"

	response "github.com/Arceister/ice-house-news/server/response"
)

type NewsHandler struct {
	service service.INewsService
}

func NewNewsHandler(service service.INewsService) handler.INewsHandler {
	return NewsHandler{
		service: service,
	}
}

func (h NewsHandler) GetNewsListHandler(w http.ResponseWriter, r *http.Request) {
	newsList, err := h.service.GetNewsListService()

	if err != nil {
		response.ErrorResponse(w, err)
		return
	}

	response.SuccessResponseWithData(w, http.StatusOK, "get news list", newsList)
}

func (h NewsHandler) GetNewsDetailHandler(w http.ResponseWriter, r *http.Request) {
	newsId := chi.URLParam(r, "newsId")
	newsDetail, err := h.service.GetNewsDetailService(newsId)

	if err != nil {
		response.ErrorResponse(w, err)
		return
	}

	response.SuccessResponseWithData(w, http.StatusOK, "get news detail", newsDetail)
}

func (h NewsHandler) AddNewNewsHandler(w http.ResponseWriter, r *http.Request) {
	var newsInput entity.NewsInputRequest

	currentUserId := r.Context().Value("JWTProps").(jwt.MapClaims)["id"].(string)
	json.NewDecoder(r.Body).Decode(&newsInput)

	err := h.service.InsertNewsService(currentUserId, newsInput)

	if err != nil {
		response.ErrorResponse(w, err)
		return
	}

	response.SuccessResponse(w, http.StatusOK, "insert new news success")
}

func (h NewsHandler) UpdateNewsHandler(w http.ResponseWriter, r *http.Request) {
	var newsInput entity.NewsInputRequest

	currentUserId := r.Context().Value("JWTProps").(jwt.MapClaims)["id"].(string)
	newsId := chi.URLParam(r, "newsId")

	json.NewDecoder(r.Body).Decode(&newsInput)

	err := h.service.UpdateNewsService(currentUserId, newsId, newsInput)

	if err != nil {
		response.ErrorResponse(w, err)
		return
	}

	response.SuccessResponse(w, http.StatusOK, "news updated")
}

func (h NewsHandler) DeleteNewsHandler(w http.ResponseWriter, r *http.Request) {
	currentUserId := r.Context().Value("JWTProps").(jwt.MapClaims)["id"].(string)
	newsId := chi.URLParam(r, "newsId")

	err := h.service.DeleteNewsService(currentUserId, newsId)

	if err != nil {
		response.ErrorResponse(w, err)
		return
	}

	response.SuccessResponse(w, http.StatusOK, "news deleted")
}
