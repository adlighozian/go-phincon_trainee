package rest

import (
	"github.com/gin-gonic/gin"
)


func ResponseData(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, map[string]interface{}{
		"status": status,
		"data"  : data,
	})
}

func ResponseError(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, map[string]interface{}{
		"status": status,
		"error":  err.Error(),
	})
}