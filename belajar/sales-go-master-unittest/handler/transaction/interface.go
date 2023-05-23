package transaction

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Handlerer interface {
	GetTransactionByNumber(w http.ResponseWriter, r *http.Request)
	CreateBulkTransactionDetail(w http.ResponseWriter, r *http.Request)
}

type GinHandlerer interface {
	GetTransactionByNumber(ctx *gin.Context)
	CreateBulkTransactionDetail(ctx *gin.Context)
}
