package main

import (
	"fmt"
	"inventory/handler"
	"net/http"
)

func main() {
	// productRepo := repository.NewProductRepository()
	// purchaseOrderRepo := repository.NewPurchaseOrderRepository(productRepo)
	// salesOrderRepo := repository.NewSalesOrderRepository()
	// inventoryHandler := handler.NewInventoryHandler(productRepo, purchaseOrderRepo, salesOrderRepo)

	// productRepo.DecodeProduct()
	// purchaseOrderRepo.DecodePurchaseOrder()

	// handler.Menu(inventoryHandler)

	Inventory := handler.NewProductHandlerHttp()
	NewServer(Inventory)
}

func NewServer(handle handler.ProductHandlerHttp) {
	// config := config.LoadConfig()
	// server
	mux := http.NewServeMux()

	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handle.ProductGet(w, r)
		}
	})

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Server running")
}
