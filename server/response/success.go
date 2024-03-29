package server

import "net/http"

type SuccessMessage struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"-"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func SuccessResponse(
	w http.ResponseWriter,
	statusCode int,
	message string,
) {
	var response SuccessMessage
	response.Success = true
	response.StatusCode = statusCode
	response.Message = message

	WriteResponse(w, int32(response.StatusCode), response)
}

func SuccessResponseWithData(
	w http.ResponseWriter,
	statusCode int,
	message string,
	data interface{},
) {
	var response SuccessMessage
	response.Success = true
	response.StatusCode = statusCode
	response.Message = message
	response.Data = data

	WriteResponse(w, int32(response.StatusCode), response)
}
