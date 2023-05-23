package transaction

import (
	"fmt"
	"net/http"
	"sales-go/helpers/gin-rest"
	"sales-go/model"
	"sales-go/usecase/transaction"
	"strconv"

	"github.com/gin-gonic/gin"
)

type gindbhttphandler struct {
	usecase transaction.TransactionUseCase
}

func NewGinDBHTTPHandler(usecase transaction.TransactionUseCase) *gindbhttphandler {
	return &gindbhttphandler{
		usecase: usecase,
	}
}

func (handler gindbhttphandler) GetTransactionByNumber(ctx *gin.Context) {
	transactionNumberStr, ok := ctx.GetQuery("transaction_id")
	if !ok {
		rest.ResponseError(ctx, http.StatusBadRequest, fmt.Errorf("transaction number should not be empty"))
		return
	}

	transactionNumber, err := strconv.Atoi(transactionNumberStr)
	if err != nil {
		rest.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := handler.usecase.GetTransactionByNumber(transactionNumber)
	if err != nil {
		rest.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}
	rest.ResponseData(ctx, http.StatusOK, res)
}

func (handler gindbhttphandler) CreateBulkTransactionDetail(ctx *gin.Context) {
	req := model.TransactionDetailBulkRequest{}

	err := ctx.ShouldBind(&req)
	if err != nil {
		rest.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	voucherCode, ok := ctx.GetQuery("voucher_code")
	if !ok {
		rest.ResponseError(ctx, http.StatusBadRequest, fmt.Errorf("voucher_id should not be empty"))
	}

	res, err := handler.usecase.CreateBulkTransactionDetail(voucherCode, req)
	if err != nil {
		if err.Error() == "name should not be empty" || err.Error() == "price should be > 0" {
			rest.ResponseError(ctx, http.StatusBadRequest, err)
			return
		} else {
			rest.ResponseError(ctx, http.StatusInternalServerError, err)
			return
		}
	}

	rest.ResponseData(ctx, http.StatusCreated, res)
}