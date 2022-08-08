package handler

import (
	"net/http"

	"github.com/Arceister/ice-house-news/handler"

	response "github.com/Arceister/ice-house-news/server/response"
	categoriesService "github.com/Arceister/ice-house-news/service/categories"
)

var _ handler.ICategoriesHandler = (*CategoriesHandler)(nil)

type CategoriesHandler struct {
	service categoriesService.CategoriesService
}

func NewCategoriesHandler(service categoriesService.CategoriesService) CategoriesHandler {
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
