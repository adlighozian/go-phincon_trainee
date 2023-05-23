package main

import (
	"inventory/controller"
	"inventory/db"
	"inventory/helper"
	"inventory/helper/middleware"
	"inventory/publisher"
	"inventory/repository/product"
	"inventory/repository/purchase"
	"inventory/repository/sales"
	"inventory/service"

	"github.com/gin-gonic/gin"
)

func main() {

	db := db.GetConnection()
	random := helper.NewRandom()
	publishPurchase := publisher.Newpurchase()
	publishSales := publisher.NewSales()

	repoProduct := product.NewProductRepository(db)
	repoPurchase := purchase.NewPurchaseRepository(db, publishPurchase, random)
	repoSales := sales.NewSalesRepository(db, publishSales, random)

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
	r.Use(middleware.Logger())

	r1 := r.Group("/inven", middleware.HeaderVerificationMiddleware)
	// routes
	r1.GET("/product", controller.ProductShow)
	r1.POST("/purchase", controller.PurchaseInput)
	r1.GET("/purchase/:order", controller.PurchaseDetail)
	r1.POST("/sales", controller.SalesInput)
	r1.GET("/sales/:order", controller.SalesDetail)

	r.GET("/auth", controller.Authentication)

	r.Run(":5000")
}
