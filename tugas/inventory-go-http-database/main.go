package main

import (
	"fmt"
	"inventory/config"
	"inventory/controller"
	"inventory/service"
	"net/http"
)

func main() {
	product := service.NewProductService()
	purchase := service.NewPurchaseService()
	// salesOrderRepo := repository.NewSalesOrderRepository(productRepo)
	Inventory := controller.NewHandlerHttp(product, purchase)
	NewServer(Inventory)
}

func NewServer(handle controller.InventoryHandlerHttp) {
	config := config.LoadConfig()
	// server
	fmt.Println("Server running...")
	mux := http.NewServeMux()

	mux.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handle.ProductShow(w, r)
		}
	})
	mux.HandleFunc("/purchase", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handle.PurchaseDetail(w, r)
		} else if r.Method == http.MethodPost {
			handle.PurchaseInput(w, r)
		}
	})
	// mux.HandleFunc("/sales", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == http.MethodGet {
	// 		handle.SalesGet(w, r)
	// 	} else if r.Method == http.MethodPost {
	// 		handle.SalesPost(w, r)
	// 	}
	// })

	server := http.Server{
		Addr:    config.JsonPort,
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
