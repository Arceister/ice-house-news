package utils

import "net/http"

func NewBadRequestError(message string) IErrorMessage {
	return &ErrorMessage{
		StatusCode:   http.StatusBadRequest,
		ErrorMessage: message,
	}
}

func NewUnauthorizedError(message string) IErrorMessage {
	return &ErrorMessage{
		StatusCode:   http.StatusUnauthorized,
		ErrorMessage: message,
	}
}

func NewNotFoundError(message string) IErrorMessage {
	return &ErrorMessage{
		StatusCode:   http.StatusNotFound,
		ErrorMessage: message,
	}
}

func NewUnprocessableEntityError(message string) IErrorMessage {
	return &ErrorMessage{
		StatusCode:   http.StatusUnprocessableEntity,
		ErrorMessage: message,
	}
}

func NewInternalServerError(message string) IErrorMessage {
	return &ErrorMessage{
		StatusCode:   http.StatusInternalServerError,
		ErrorMessage: message,
	}
}
