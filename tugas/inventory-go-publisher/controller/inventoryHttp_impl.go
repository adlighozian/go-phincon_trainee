package controller

import (
	"inventory/helper"
	"inventory/helper/middleware"
	"inventory/model"
	"inventory/service"
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

	keyss := c.GetHeader("key")

	err := helper.ComparePassword(keyss, "phincon")
	if err != nil {
		c.JSON(http.StatusOK, map[string]string{
			"message": "Gagal",
		})
	} else {
		c.JSON(http.StatusOK, map[string]string{
			"message": "Berhasil",
		})
	}

}

// product
func (handler *handlerHttp) ProductShow(c *gin.Context) {
	result, _ := handler.product.ShowProduct()
	c.JSON(result.Status, result)
}

// purchase
func (handler *handlerHttp) PurchaseInput(c *gin.Context) {
	var data []model.ReqPurchase

	err := c.ShouldBindJSON(&data)
	middleware.FailError(err, "baris satu")

	result, _ := handler.purchase.InputPurchase(data)

	c.JSON(result.Status, result)
}

func (handler *handlerHttp) PurchaseDetail(c *gin.Context) {
	order := c.Param("order")

	result, _ := handler.purchase.DetailPurchase(order)
	c.JSON(result.Status, result)
}

// sales
func (handler *handlerHttp) SalesInput(c *gin.Context) {
	var data []model.ReqSales

	err := c.ShouldBindJSON(&data)
	middleware.FailError(err, "baris satu")

	result, _ := handler.sales.InputSales(data)

	c.JSON(result.Status, result)
}

func (handler *handlerHttp) SalesDetail(c *gin.Context) {
	order := c.Param("order")

	result, _ := handler.sales.DetailSales(order)
	c.JSON(result.Status, result)
}
