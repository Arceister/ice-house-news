package server

import (
	"net/http"

	error "github.com/Arceister/ice-house-news/utils/error"
)

type ErrorMessage struct {
	Success      bool        `json:"success"`
	StatusCode   int         `json:"-"`
	ErrorMessage string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
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
