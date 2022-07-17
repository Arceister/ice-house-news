package handler

import "github.com/Arceister/ice-house-news/service"

type NewsHandler struct {
	service service.UsersService
}

func NewNewsHandler(service service.UsersService) NewsHandler {
	return NewsHandler{
		service: service,
	}
}
