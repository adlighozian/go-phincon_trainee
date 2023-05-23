package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sales-go/helpers/gin-rest"
	"sales-go/model"
	"sales-go/usecase/product"
)

type gindbhttphandler struct {
	usecase product.ProductUseCase
}

func NewGinDBHTTPHandler(usecase product.ProductUseCase) *gindbhttphandler {
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
	req := []model.ProductRequest{}

	err := ctx.ShouldBind(&req)
	if err != nil {
		rest.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	} else if req == nil {
		rest.ResponseError(ctx, http.StatusBadRequest, err)
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