package utils

import "net/http"

func NewBadRequestError(message string) IErrorMessage {
	return &ErrorMessage{
		StatusCode:    http.StatusBadRequest,
		ErrorResponse: "bad_request",
		ErrorMessage:  message,
	}
}

func NewUnauthorizedError(message string) IErrorMessage {
	return &ErrorMessage{
		StatusCode:    http.StatusUnauthorized,
		ErrorResponse: "unauthorized",
		ErrorMessage:  message,
	}
}

func NewNotFoundError(message string) IErrorMessage {
	return &ErrorMessage{
		StatusCode:    http.StatusNotFound,
		ErrorResponse: "not_found",
		ErrorMessage:  message,
	}
}

func NewUnprocessableEntityError(message string) IErrorMessage {
	return &ErrorMessage{
		StatusCode:    http.StatusUnprocessableEntity,
		ErrorResponse: "unprocessable_entity",
		ErrorMessage:  message,
	}
}

func NewInternalServerError(message string) IErrorMessage {
	return &ErrorMessage{
		StatusCode:    http.StatusInternalServerError,
		ErrorResponse: "internal_server_error",
		ErrorMessage:  message,
	}
}
