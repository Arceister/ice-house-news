package server

import (
	"encoding/json"
	"net/http"
)

type ResponseStruct struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ResponseWithData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseJSON(
	w http.ResponseWriter,
	statusCode int32,
	success bool,
	message string,
) {
	var response ResponseStruct
	response.Success = success
	response.Message = message

	writeResponse(w, statusCode, response)
}

func ResponseJSONData(
	w http.ResponseWriter,
	statusCode int32,
	success bool,
	message string,
	data interface{},
) {

	var response ResponseWithData
	response.Success = success
	response.Message = message
	response.Data = data

	writeResponse(w, statusCode, response)
}

func writeResponse(
	w http.ResponseWriter,
	statusCode int32,
	response interface{},
) {
	res, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(statusCode))
	w.Write(res)
}
