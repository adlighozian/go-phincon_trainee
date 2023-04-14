package controller

import "net/http"

type InventoryHandlerHttp interface {
	ProductShow(w http.ResponseWriter, r *http.Request)
	PurchaseDetail(w http.ResponseWriter, r *http.Request)
	PurchaseInput(w http.ResponseWriter, r *http.Request)
	// SalesGet(w http.ResponseWriter, r *http.Request)
	// SalesPost(w http.ResponseWriter, r *http.Request)
}
