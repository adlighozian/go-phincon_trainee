package controller

import "net/http"

type InventoryHandlerHttp interface {
	ProductShow(w http.ResponseWriter, r *http.Request)
	PurchaseDetail(w http.ResponseWriter, r *http.Request)
	PurchaseInput(w http.ResponseWriter, r *http.Request)
	SalesDetail(w http.ResponseWriter, r *http.Request)
	SalesInput(w http.ResponseWriter, r *http.Request)
}
