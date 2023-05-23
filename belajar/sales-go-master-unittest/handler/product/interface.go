package product

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Handlerer interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type GinHandlerer interface {
	GetList(ctx *gin.Context)
	Create(ctx *gin.Context)
}
