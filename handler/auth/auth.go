package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/server"
	users "github.com/Arceister/ice-house-news/service/users"
)

type AuthHandler struct {
	userService users.UsersService
}

func NewAuthHandler(userService users.UsersService) AuthHandler {
	return AuthHandler{
		userService: userService,
	}
}

func (c AuthHandler) UserSignInHandler(w http.ResponseWriter, r *http.Request) {
	var userInput entity.UserSignInRequest
	json.NewDecoder(r.Body).Decode(&userInput)

	token, err := c.userService.SignInService(userInput)
	if err != nil && (err.Error() == "user not found" || err.Error() == "wrong password") {
		server.ResponseJSON(w, http.StatusUnauthorized, false, err.Error())
		return
	} else if err != nil {
		server.ResponseJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	var userToken entity.UserToken
	userToken.Token = *token

	server.ResponseJSONData(w, http.StatusOK, true, "Login successful!", userToken)
}