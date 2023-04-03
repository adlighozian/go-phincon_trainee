package handler

import "net/http"

type ContactHandlerHttp interface {
	HandlerGet(write http.ResponseWriter, request *http.Request)
	HandlerPost(write http.ResponseWriter, request *http.Request)
	HandlerUpdate(write http.ResponseWriter, request *http.Request)
	HandlerDelete(write http.ResponseWriter, request *http.Request)
}

type ProductHandlerHttp interface {
	ProductGet(write http.ResponseWriter, request *http.Request)
}
