package server

import (
	"net/http"

	error "github.com/Arceister/ice-house-news/utils/error"
)

type ErrorMessage struct {
	MessageStruct
	StatusCode   int    `json:"-"`
	ErrorMessage string `json:"message"`
}

func ErrorResponse(
	w http.ResponseWriter,
	errorMessage error.IErrorMessage,
) {
	var response ErrorMessage
	response.Success = false
	response.StatusCode = errorMessage.Status()
	response.ErrorMessage = errorMessage.Message()

	WriteResponse(w, int32(response.StatusCode), response)
}
