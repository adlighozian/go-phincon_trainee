package main

import (
	"inventory/controller"
	"inventory/db"
	"inventory/helper/middleware"
	"inventory/repository/product"
	"inventory/repository/purchase"
	"inventory/repository/sales"
	"inventory/service"

	"github.com/gin-gonic/gin"
)

func main() {

	db := db.GetConnection()

	repoProduct := product.NewProductRepository(db)
	repoPurchase := purchase.NewPurchaseRepository(db)
	repoSales := sales.NewSalesRepository(db)

	purchase := service.NewPurchaseService(repoPurchase)
	product := service.NewProductService(repoProduct)
	sales := service.NewSalesService(repoSales)

	Inventory := controller.NewHandlerHttp(product, purchase, sales)
	NewServer(Inventory)
	// consumer.ConPurchase()
}

func NewServer(controller controller.InventoryHandlerHttp) {
	// server
	r := gin.Default()

	// middleware
	r.Use(middleware.Logger(), middleware.CheckAuth())

	// routes
	r.GET("/product", controller.ProductShow)
	r.GET("/auth", controller.Authentication)
	// purchase
	r.POST("/purchase", controller.PurchaseInput)
	r.GET("/purchase/:order", controller.PurchaseDetail)
	// sales
	r.POST("/sales", controller.SalesInput)
	r.GET("/sales/:order", controller.SalesDetail)

	r.Run(":5000")
}
