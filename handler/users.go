package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/server"
	service "github.com/Arceister/ice-house-news/service"
	"github.com/go-chi/chi/v5"
)

type UsersHandler struct {
	service service.UsersService
}

func NewUsersHandler(service service.UsersService) UsersHandler {
	return UsersHandler{
		service: service,
	}
}

func (c UsersHandler) GetOneUserHandler(w http.ResponseWriter, r *http.Request) {
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

func (c UsersHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userData entity.User
	json.NewDecoder(r.Body).Decode(&userData)

	err := c.service.CreateUserService(userData)

	if err != nil {
		server.ResponseJSON(w, 500, false, err.Error())
		return
	}

	server.ResponseJSON(w, 200, true, "User create success")
}

func (c UsersHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
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

func (c UsersHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
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