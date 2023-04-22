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
	sales := service.NewSalesService()

	Inventory := controller.NewHandlerHttp(product, purchase, sales)
	NewServer(Inventory)
}

func NewServer(handle controller.InventoryHandlerHttp) {
	config := config.LoadConfig()
	// server
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
	mux.HandleFunc("/sales", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handle.SalesDetail(w, r)
		} else if r.Method == http.MethodPost {
			handle.SalesInput(w, r)
		}
	})

	// middleware := middleware.Use(middleware.LoggingHandler(mux))

	server := http.Server{
		Addr:    config.JsonPort,
		Handler: mux,
		// Handler: middleware[0],
	}
	fmt.Println("Server running on", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
