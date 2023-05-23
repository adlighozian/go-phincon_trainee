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
	publish := publisher.NewPublisher()
	random := helper.NewRandom()

	repoProduct := product.NewProductRepository(db)
	repoPurchase := purchase.NewPurchaseRepository(db, publish, random)
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
	r.Use(middleware.Logger())

	r1 := r.Group("/inven")
	r1.Use(middleware.CheckAuth())
	// routes
	r1.GET("/product", controller.ProductShow)
	r1.POST("/purchase", controller.PurchaseInput)
	r1.GET("/purchase/:order", controller.PurchaseDetail)
	r1.POST("/sales", controller.SalesInput)
	r1.GET("/sales/:order", controller.SalesDetail)

	r.GET("/auth", controller.Authentication)

	r.Run(":5000")
}
