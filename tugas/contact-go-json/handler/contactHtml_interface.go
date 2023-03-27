package handler

import "net/http"

type ContactHandlerHttp interface {
	HandlerGet(write http.ResponseWriter, request *http.Request)
	HandlerPost(write http.ResponseWriter, request *http.Request)
}
