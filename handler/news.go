package handler

import (
	"net/http"

	"github.com/Arceister/ice-house-news/server"
	"github.com/Arceister/ice-house-news/service"
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
