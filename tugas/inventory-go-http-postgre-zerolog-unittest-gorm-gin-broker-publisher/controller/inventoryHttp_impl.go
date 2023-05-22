package controller

import (
	"inventory/helper/middleware"
	"inventory/model"
	"inventory/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerHttp struct {
	product  service.ProductService
	purchase service.PurchaseService
	sales    service.SalesService
}

func NewHandlerHttp(productService service.ProductService, purchaseService service.PurchaseService, salesService service.SalesService) InventoryHandlerHttp {
	return &handlerHttp{
		product:  productService,
		purchase: purchaseService,
		sales:    salesService,
	}
}

func (handler *handlerHttp) Authentication(c *gin.Context) {
	// token := "$2a$12$QUCKWQQYVLsvn663mnp9XuVnz7yei2qTsj/DhZkDDfH5qnf6Cfpti"
	// c.Header("auth", token)

	req := c.Request.Header.Get("auth")
	c.JSON(http.StatusOK, req)
}

// product
func (handler *handlerHttp) ProductShow(c *gin.Context) {
	log.Println("product controller")
	result, _ := handler.product.ShowProduct()

	c.JSON(result.Status, result)
}

// purchase
func (handler *handlerHttp) PurchaseInput(c *gin.Context) {
	log.Println("purchase controller")

	var data []model.ReqPurchase

	err := c.ShouldBindJSON(&data)
	middleware.FailError(err, "baris satu")

	result, _ := handler.purchase.InputPurchase(data)

	c.JSON(result.Status, result)
}

func (handler *handlerHttp) PurchaseDetail(c *gin.Context) {
	log.Println("purchase controller")

	order := c.Param("order")

	result, _ := handler.purchase.DetailPurchase(order)
	c.JSON(result.Status, result)
}

// sales
func (handler *handlerHttp) SalesInput(c *gin.Context) {
	log.Println("sales controller")

	var data []model.ReqSales

	err := c.ShouldBindJSON(&data)
	middleware.FailError(err, "baris satu")

	result, _ := handler.sales.InputSales(data)

	c.JSON(result.Status, result)
}

func (handler *handlerHttp) SalesDetail(c *gin.Context) {
	log.Println("sales controller")

	order := c.Param("order")

	result, _ := handler.sales.DetailSales(order)
	c.JSON(result.Status, result)
}
