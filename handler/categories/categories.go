package handler

import (
	"net/http"

	"github.com/Arceister/ice-house-news/handler"
	"github.com/Arceister/ice-house-news/server"
	"github.com/Arceister/ice-house-news/service"
)

type CategoriesHandler struct {
	service service.ICategoriesService
}

func NewCategoriesHandler(service service.ICategoriesService) handler.ICategoriesHandler {
	return CategoriesHandler{
		service: service,
	}
}

func (h CategoriesHandler) GetAllNewsCategoryHandler(w http.ResponseWriter, r *http.Request) {
	newsCategories, err := h.service.GetAllNewsCategoryService()

	if err != nil {
		server.ResponseJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	server.ResponseJSONData(w, http.StatusOK, true, "get all categories", newsCategories)
}
