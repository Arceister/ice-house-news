package handler

import (
	"net/http"

	"github.com/Arceister/ice-house-news/handler"
	"github.com/Arceister/ice-house-news/service"

	response "github.com/Arceister/ice-house-news/server/response"
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
		response.ErrorResponse(w, err)
		return
	}

	response.SuccessResponseWithData(w, http.StatusOK, "get all categories", newsCategories)
}
