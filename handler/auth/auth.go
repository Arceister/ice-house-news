package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/handler"
	"github.com/Arceister/ice-house-news/server"
	"github.com/Arceister/ice-house-news/service"
)

type AuthHandler struct {
	userService service.IUsersService
}

func NewAuthHandler(userService service.IUsersService) handler.IAuthHandler {
	return AuthHandler{
		userService: userService,
	}
}

func (c AuthHandler) UserSignInHandler(w http.ResponseWriter, r *http.Request) {
	var userInput entity.UserSignInRequest
	json.NewDecoder(r.Body).Decode(&userInput)

	signInSchema, err := c.userService.SignInService(userInput)
	if err != nil && (err.Error() == "user not found" || err.Error() == "wrong password") {
		server.ResponseJSON(w, http.StatusUnauthorized, false, err.Error())
		return
	} else if err != nil {
		server.ResponseJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	server.ResponseJSONData(w, http.StatusOK, true, "Login successful!", signInSchema)
}
