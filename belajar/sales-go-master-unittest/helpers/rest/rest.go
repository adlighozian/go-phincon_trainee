package rest

import (
	"encoding/json"
	"net/http"

	logger "sales-go/helpers/logging"
)

type RestError struct {
	Status  int			`json:"status"`
	Message string		`json:"message"`
	Error	string		`json:"error"`
}

func ResponseError(w http.ResponseWriter, r *http.Request, status int, err error) {
	logger.Errorf(err, r)

	response := RestError{
		Status:  status,
		Message: http.StatusText(status),
		Error:	 err.Error(),
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(status)
	w.Write(jsonData)
}

type RestData struct {
	Status  int			`json:"status"`
	Message string		`json:"message"`
	Data	interface{}	`json:"data"`
}

func ResponseData(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	response := RestData{
		Status:  status,
		Message: http.StatusText(status),
		Data:	 data,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(status)
	w.Write(jsonData)
}
