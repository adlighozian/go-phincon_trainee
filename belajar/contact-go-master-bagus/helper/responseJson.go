package helper

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewJsonResponse(w http.ResponseWriter, code int, message string, data interface{}) error {
	res := JsonResponse{
		Status:  code,
		Message: message,
		Data:    data,
	}

	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(&res)
}
