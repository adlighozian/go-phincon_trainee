//go:generate mockery --output=../mocks --name ContactHTTPHandler
package handler

import "net/http"

type ContactHTTPHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Detail(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
