package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/handler"
	response "github.com/Arceister/ice-house-news/server/response"
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
	if err != nil {
		response.ErrorResponse(w, err)
		return
	}

	response.SuccessResponseWithData(w, 200, "Login successful", signInSchema)
}
