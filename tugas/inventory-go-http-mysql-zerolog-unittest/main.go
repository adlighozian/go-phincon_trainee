package main

import (
	"fmt"
	"inventory/config"
	"inventory/controller"
	"inventory/middleware"
	"inventory/service"
	"net/http"
)

func main() {
	config := config.LoadConfig()

	inventoryServ := createInventoryService(config)

	product := service.NewProductService()
	purchase := service.NewPurchaseService()
	sales := service.NewSalesService()

	Inventory := controller.NewHandlerHttp(product, purchase, sales)
	NewServer(Inventory)
}

func createInventoryService(cfg *config.Config) service {

}

func NewServer(handle controller.InventoryHandlerHttp) {
	config := config.LoadConfig()
	// server
	mux := http.NewServeMux()
	mux.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		handle.ProductShow(w, r)
	})
	mux.HandleFunc("/purchase/detail", func(w http.ResponseWriter, r *http.Request) {
		handle.PurchaseDetail(w, r)
	})
	mux.HandleFunc("/purchase/input", func(w http.ResponseWriter, r *http.Request) {
		handle.PurchaseInput(w, r)
	})
	mux.HandleFunc("/sales/detail", func(w http.ResponseWriter, r *http.Request) {
		handle.SalesDetail(w, r)
	})
	mux.HandleFunc("/sales/input", func(w http.ResponseWriter, r *http.Request) {
		handle.SalesInput(w, r)
	})

	middleware := middleware.Use(middleware.LoggingHandler(mux))

	server := http.Server{
		Addr: config.JsonPort,
		// Handler: mux,
		Handler: middleware[0],
	}
	fmt.Println("Server running on", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
