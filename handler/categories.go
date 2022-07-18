package handler

import "github.com/Arceister/ice-house-news/service"

type CategoriesHandler struct {
	service service.CategoriesService
}

func NewCategoriesHandler(service service.CategoriesService) CategoriesHandler {
	return CategoriesHandler{
		service: service,
	}
}
