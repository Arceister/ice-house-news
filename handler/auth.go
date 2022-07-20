package handler

import "github.com/Arceister/ice-house-news/service"

type AuthHandler struct {
	userService service.UsersService
}

func NewAuthHandler(userService service.UsersService) AuthHandler {
	return AuthHandler{
		userService: userService,
	}
}
