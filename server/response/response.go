package server

import (
	"encoding/json"
	"net/http"
)

type MessageStruct struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteResponse(
	w http.ResponseWriter,
	statusCode int32,
	response interface{},
) {
	res, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(statusCode))
	w.Write(res)
}
