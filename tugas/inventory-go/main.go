package main

import (
	"inventory/handler"
	"inventory/repository"
)

func main() {
	productRepo := repository.NewProductRepository()
	purchaseOrderRepo := repository.NewPurchaseOrderRepository()
	salesOrderRepo := repository.NewSalesOrderRepository()
	inventoryHandler := handler.NewInventoryHandler(productRepo, purchaseOrderRepo, salesOrderRepo)

	handler.Menu(inventoryHandler)
}
