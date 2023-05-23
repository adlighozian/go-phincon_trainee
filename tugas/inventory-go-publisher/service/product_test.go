package service

import (
	"errors"
	"inventory/mocks"
	"inventory/model"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceProduct(t *testing.T) {
	t.Run("get-product-success", func(t *testing.T) {
		mock := mocks.NewServProductMock()
		serv := NewProductService(mock)

		mock.On("ShowProduct").Return([]model.Product{
			{
				Id:    1,
				Name:  "laptop",
				Price: 15000,
				Stock: 10,
			},
		}, nil)

		res, err := serv.ShowProduct()
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
		require.Equal(t, "OK", res.Message)
	})

	t.Run("get-product-fail", func(t *testing.T) {
		mock := mocks.NewServProductMock()
		serv := NewProductService(mock)
		mock.On("ShowProduct").Return([]model.Product{}, errors.New("error"))
		res, err := serv.ShowProduct()
		require.Error(t, err)
		require.Equal(t, http.StatusBadGateway, res.Status)
		require.Equal(t, "Bad Gateway", res.Message)
	})

}
