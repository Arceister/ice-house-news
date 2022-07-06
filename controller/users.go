package controller

import (
	"net/http"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/server"
	service "github.com/Arceister/ice-house-news/service"
	"github.com/go-chi/chi/v5"
)

type UsersController struct {
	service service.UsersService
}

func NewUsersController(service service.UsersService) UsersController {
	return UsersController{
		service: service,
	}
}

func (c UsersController) GetOneUserController(w http.ResponseWriter, r *http.Request) {
	userUniqueId := chi.URLParam(r, "uuid")
	result, err := c.service.GetOneUserService(userUniqueId)

	if result == (entity.User{}) {
		server.ResponseJSON(w, 404, false, "user not found")
		return
	}

	if err != nil {
		server.ResponseJSON(w, 500, false, err.Error())
		return
	}

	server.ResponseJSONData(w, 200, true, "get one user", result)
}
