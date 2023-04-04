package main

import (
	"fmt"
	"inventory/config"
	"inventory/handler"
	"inventory/repository"
	"net/http"
)

func main() {
	productRepo := repository.NewProductRepository()
	purchaseOrderRepo := repository.NewPurchaseOrderRepository(productRepo)
	salesOrderRepo := repository.NewSalesOrderRepository(productRepo)
	Inventory := handler.NewHandlerHttp(productRepo, purchaseOrderRepo, salesOrderRepo)
	NewServer(Inventory)
}

func NewServer(handle handler.InventoryHandlerHttp) {
	config := config.LoadConfig()
	// server
	mux := http.NewServeMux()

	mux.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handle.ProductGet(w, r)
		}
	})
	mux.HandleFunc("/purchase", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handle.PurchaseGet(w, r)
		} else if r.Method == http.MethodPost {
			handle.PurchasePost(w, r)
		}
	})
	mux.HandleFunc("/sales", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handle.SalesGet(w, r)
		} else if r.Method == http.MethodPost {
			handle.SalesPost(w, r)
		}
	})

	server := http.Server{
		Addr:    config.Port,
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Server running")
}
