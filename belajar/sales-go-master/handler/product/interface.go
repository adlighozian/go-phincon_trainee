package product

import (
	"net/http"
)

type Handlerer interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}
