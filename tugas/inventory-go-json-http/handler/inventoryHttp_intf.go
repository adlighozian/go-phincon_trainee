package handler

import "net/http"

type InventoryHandlerHttp interface {
	ProductGet(w http.ResponseWriter, r *http.Request)
	PurchaseGet(w http.ResponseWriter, r *http.Request)
	PurchasePost(w http.ResponseWriter, r *http.Request) 
	SalesGet(w http.ResponseWriter, r *http.Request)
	SalesPost(w http.ResponseWriter, r *http.Request)
}
