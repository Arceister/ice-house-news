package controller

import (
	"encoding/json"
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

	if err != nil && err.Error() == "user not found" {
		server.ResponseJSON(w, 404, false, err.Error())
		return
	} else if err != nil {
		server.ResponseJSON(w, 500, false, err.Error())
		return
	}

	server.ResponseJSONData(w, 200, true, "get one user", result)
}

func (c UsersController) CreateUserController(w http.ResponseWriter, r *http.Request) {
	var userData entity.User
	json.NewDecoder(r.Body).Decode(&userData)

	err := c.service.CreateUserService(userData)

	if err != nil {
		server.ResponseJSON(w, 500, false, err.Error())
		return
	}

	server.ResponseJSON(w, 200, true, "User create success")
}

func (c UsersController) UpdateUserController(w http.ResponseWriter, r *http.Request) {
	userUniqueId := chi.URLParam(r, "uuid")

	var userData entity.User
	json.NewDecoder(r.Body).Decode(&userData)

	err := c.service.UpdateUserService(userUniqueId, userData)

	if err != nil && err.Error() == "user not found" {
		server.ResponseJSON(w, 404, false, err.Error())
		return
	} else if err != nil {
		server.ResponseJSON(w, 500, false, err.Error())
		return
	}

	server.ResponseJSON(w, 200, true, "User update success")
}

func (c UsersController) DeleteUserController(w http.ResponseWriter, r *http.Request) {
	userUniqueId := chi.URLParam(r, "uuid")

	err := c.service.DeleteUserService(userUniqueId)

	if err != nil && err.Error() == "user not found" {
		server.ResponseJSON(w, 404, false, err.Error())
		return
	} else if err != nil {
		server.ResponseJSON(w, 500, false, err.Error())
		return
	}

	server.ResponseJSON(w, 200, true, "User deleted")
}
