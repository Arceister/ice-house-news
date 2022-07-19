package controller

import service "github.com/Arceister/ice-house-news/service"

type UsersController struct {
	service service.UsersService
}

func NewUsersController(service service.UsersService) UsersController {
	return UsersController{
		service: service,
	}
}
