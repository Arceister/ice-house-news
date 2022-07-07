package controller

import "github.com/Arceister/ice-house-news/service"

type AuthController struct {
	userService service.UsersService
}

func NewAuthController(userService service.UsersService) AuthController {
	return AuthController{
		userService: userService,
	}
}
