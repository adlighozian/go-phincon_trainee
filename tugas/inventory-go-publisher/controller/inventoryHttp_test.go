package controller

import (
	"inventory/mocks"
	"inventory/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestController(t *testing.T) {

	t.Run("authentication-success", func(t *testing.T) {
		// mock := mocks.NewControllerMock()
		// cnt := NewHandlerHttp(, mock, mock)

		// mock.On("ComparePassword").Return(nil)

		// cnt.Authentication(c * gin.Contex)
	})

	t.Run("product-show-success", func(t *testing.T) {
		productMock := mocks.NewProductControllerMock()
		purchaseMock := mocks.NewPurchaseControllerMock()
		salesMock := mocks.NewSalesControllerMock()
		cnt := NewHandlerHttp(productMock, purchaseMock, salesMock)

		productMock.On("ShowProduct").Return(model.InventoryResponse{
			Status:  http.StatusOK,
			Message: "OK",
			Data:    model.Product{},
		}, nil)

		request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/product", nil)
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		c.Request = request

		cnt.ProductShow(c)
		response := recorder.Result()
		require.Equal(t, http.StatusOK, response.StatusCode)
	})
}
