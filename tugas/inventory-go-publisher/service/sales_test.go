package service

import (
	"errors"
	"inventory/mocks"
	"inventory/model"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceSales(t *testing.T) {

	t.Run("input-sales-success", func(t *testing.T) {
		mock := mocks.NewServSalesMock()
		serv := NewSalesService(mock)

		req := []model.ReqSales{
			{
				Item:  "hp",
				Price: 15000,
				From:  "bagas",
				Total: 10,
			},
		}

		mock.On("InputSales", req).Return([]model.SalesDetail{
			{
				Id:       1,
				Sales_id: 1,
				Item:     "hp",
				Price:    15000,
				Quantity: 10,
				Total:    10,
				Sales: model.Sales{
					Id:          1,
					OrderNumber: "1221",
					From:        "bagas",
					Total:       10,
				},
			},
		}, nil)

		res, err := serv.InputSales(req)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
		require.Equal(t, "OK", res.Message)
	})

	t.Run("input-sales-fail", func(t *testing.T) {
		mock := mocks.NewServSalesMock()
		serv := NewSalesService(mock)

		req := []model.ReqSales{
			{
				Item:  "hp",
				Price: 15000,
				From:  "bagas",
				Total: 10,
			},
		}

		mock.On("InputSales", req).Return([]model.SalesDetail{}, errors.New("error"))

		res, err := serv.InputSales(req)

		require.Error(t, err)
		require.Equal(t, http.StatusBadRequest, res.Status)
		require.Equal(t, "Bad Request", res.Message)
	})

	t.Run("detail-sales-success", func(t *testing.T) {
		mock := mocks.NewServSalesMock()
		serv := NewSalesService(mock)

		req := "1221"

		mock.On("DetailSales", req).Return(model.SalesDetail{
			Id:       1,
			Sales_id: 1,
			Item:     "hp",
			Price:    15000,
			Quantity: 10,
			Total:    10,
			Sales: model.Sales{
				Id:          1,
				OrderNumber: "1221",
				From:        "bagas",
				Total:       10,
			},
		}, nil)

		res, err := serv.DetailSales(req)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
		require.Equal(t, "OK", res.Message)

	})

	t.Run("detail-sales-fail", func(t *testing.T) {
		mock := mocks.NewServSalesMock()
		serv := NewSalesService(mock)

		req := "1221"

		mock.On("DetailSales", req).Return(model.SalesDetail{}, errors.New("error"))

		res, err := serv.DetailSales(req)

		require.Error(t, err)
		require.Equal(t, http.StatusBadRequest, res.Status)
		require.Equal(t, "Bad Request", res.Message)

	})
}
