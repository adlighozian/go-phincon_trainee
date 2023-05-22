package controller

import "github.com/gin-gonic/gin"

type InventoryHandlerHttp interface {
	Authentication(c *gin.Context)
	ProductShow(c *gin.Context)
	PurchaseInput(c *gin.Context)
	PurchaseDetail(c *gin.Context)
	SalesInput(c *gin.Context)
	SalesDetail(c *gin.Context)
}
