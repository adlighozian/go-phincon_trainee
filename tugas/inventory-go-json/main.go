package main

import (
	"inventory/handler"
	"inventory/repository"
)

func main() {
	productRepo := repository.NewProductRepository()
	purchaseOrderRepo := repository.NewPurchaseOrderRepository(productRepo)
	salesOrderRepo := repository.NewSalesOrderRepository(productRepo)
	inventoryHandler := handler.NewInventoryHandler(productRepo, purchaseOrderRepo, salesOrderRepo)

	productRepo.DecodeProduct()
	purchaseOrderRepo.DecodePurchaseOrder()

	handler.Menu(inventoryHandler)

}
