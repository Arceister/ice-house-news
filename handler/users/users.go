package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Arceister/ice-house-news/entity"
	response "github.com/Arceister/ice-house-news/server/response"
	usersService "github.com/Arceister/ice-house-news/service/users"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
)

type UsersHandler struct {
	service usersService.UsersService
}

func NewUsersHandler(service usersService.UsersService) UsersHandler {
	return UsersHandler{
		service: service,
	}
}

func (c UsersHandler) GetOwnProfile(w http.ResponseWriter, r *http.Request) {
	currentUserId := r.Context().Value("JWTProps").(jwt.MapClaims)["id"].(string)
	result, err := c.service.GetOneUserService(currentUserId)

	if err != nil {
		response.ErrorResponse(w, err)
		return
	}

	response.SuccessResponseWithData(w, 200, "get profile", result)
}

func (c UsersHandler) GetOneUserHandler(w http.ResponseWriter, r *http.Request) {
	userUniqueId := chi.URLParam(r, "uuid")
	result, err := c.service.GetOneUserService(userUniqueId)

	if err != nil {
		response.ErrorResponse(w, err)
		return
	}

	response.SuccessResponseWithData(w, 200, "get one user", result)
}

func (c UsersHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userData entity.User
	json.NewDecoder(r.Body).Decode(&userData)

	err := c.service.CreateUserService(userData)

	if err != nil {
		response.ErrorResponse(w, err)
		return
	}

	response.SuccessResponse(w, 200, "create user success")
}

func (c UsersHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	userUniqueId := chi.URLParam(r, "uuid")

	var userData entity.User
	json.NewDecoder(r.Body).Decode(&userData)

	err := c.service.UpdateUserService(userUniqueId, userData)

	if err != nil {
		response.ErrorResponse(w, err)
		return
	}

	response.SuccessResponse(w, 200, "update user success")
}

func (c UsersHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userUniqueId := chi.URLParam(r, "uuid")

	err := c.service.DeleteUserService(userUniqueId)

	if err != nil {
		response.ErrorResponse(w, err)
		return
	}

	response.SuccessResponse(w, 200, "delete user success")
}
