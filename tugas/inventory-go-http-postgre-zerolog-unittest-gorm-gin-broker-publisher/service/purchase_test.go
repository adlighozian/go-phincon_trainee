package service

import (
	"errors"
	"inventory/mocks"
	"inventory/model"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServicePurchase(t *testing.T) {

	t.Run("input-purchase-success", func(t *testing.T) {
		mock := mocks.NewServPurchaseMock()
		serv := NewPurchaseService(mock)

		req := []model.ReqPurchase{
			{
				Item:  "hp",
				Price: 15000,
				From:  "bagas",
				Total: 10,
			},
		}

		mock.On("InputPurchase", req).Return([]model.PurchaseDetail{
			{
				Id:          1,
				Purchase_id: 1,
				Item:        "hp",
				Price:       15000,
				Quantity:    10,
				Total:       10,
				Purchase: model.Purchase{
					Id:          1,
					OrderNumber: "1221",
					From:        "bagas",
					Total:       10,
				},
			},
		}, nil)

		res, err := serv.InputPurchase(req)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
		require.Equal(t, "OK", res.Message)
	})

	t.Run("input-purchase-fail", func(t *testing.T) {
		mock := mocks.NewServPurchaseMock()
		serv := NewPurchaseService(mock)

		req := []model.ReqPurchase{
			{
				Item:  "hp",
				Price: 15000,
				From:  "bagas",
				Total: 10,
			},
		}

		mock.On("InputPurchase", req).Return([]model.PurchaseDetail{}, errors.New("error"))

		res, err := serv.InputPurchase(req)

		require.Error(t, err)
		require.Equal(t, http.StatusBadRequest, res.Status)
		require.Equal(t, "Bad Request", res.Message)
	})

	t.Run("detail-purchase-success", func(t *testing.T) {
		mock := mocks.NewServPurchaseMock()
		serv := NewPurchaseService(mock)

		req := "1221"

		mock.On("DetailPurchase", req).Return(model.PurchaseDetail{
			Id:          1,
			Purchase_id: 1,
			Item:        "hp",
			Price:       15000,
			Quantity:    10,
			Total:       10,
			Purchase: model.Purchase{
				Id:          1,
				OrderNumber: "1221",
				From:        "bagas",
				Total:       10,
			},
		}, nil)

		res, err := serv.DetailPurchase(req)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
		require.Equal(t, "OK", res.Message)
	})
	t.Run("detail-purchase-fail", func(t *testing.T) {
		mock := mocks.NewServPurchaseMock()
		serv := NewPurchaseService(mock)

		req := "1221"

		mock.On("DetailPurchase", req).Return(model.PurchaseDetail{}, errors.New("error"))

		res, err := serv.DetailPurchase(req)

		require.Error(t, err)
		require.Equal(t, http.StatusBadRequest, res.Status)
		require.Equal(t, "Bad Request", res.Message)
	})
}
