package voucher

import (
	"net/http"
	"sales-go/helpers/gin-rest"
	"sales-go/model"
	"sales-go/usecase/voucher"

	"github.com/gin-gonic/gin"
)

type gindbhttphandler struct {
	usecase voucher.VoucherUseCase
}

func NewGinDBHTTPHandler(usecase voucher.VoucherUseCase) *gindbhttphandler {
	return &gindbhttphandler{
		usecase: usecase,
	}
}

func (handler gindbhttphandler) GetList(ctx *gin.Context) {
	res, err := handler.usecase.GetList()
	if err != nil {
		rest.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}
	rest.ResponseData(ctx, http.StatusOK, res)
}

func (handler gindbhttphandler) Create(ctx *gin.Context) {
	req := []model.VoucherRequest{}

	err := ctx.ShouldBind(&req)
	if err != nil {
		rest.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	res, err := handler.usecase.Create(req)
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