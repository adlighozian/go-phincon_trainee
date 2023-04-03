package handler

import "net/http"

type ProductHandlerHttp interface {
	ProductGet(write http.ResponseWriter, request *http.Request)
}
