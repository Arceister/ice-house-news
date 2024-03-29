package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/handler"
	response "github.com/Arceister/ice-house-news/server/response"
	"github.com/Arceister/ice-house-news/service"
	"github.com/golang-jwt/jwt/v4"
)

var _ handler.IAuthHandler = (*AuthHandler)(nil)

type AuthHandler struct {
	userService service.IUsersService
}

func NewAuthHandler(userService service.IUsersService) *AuthHandler {
	return &AuthHandler{
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

func (c AuthHandler) ExtendTokenHandler(w http.ResponseWriter, r *http.Request) {
	currentUserId := r.Context().Value("JWTProps").(jwt.MapClaims)["id"].(string)

	signInSchema, err := c.userService.ExtendToken(currentUserId)
	if err != nil {
		response.ErrorResponse(w, err)
		return
	}

	response.SuccessResponseWithData(w, 200, "Login successful", signInSchema)
}
