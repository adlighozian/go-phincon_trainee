package service

import (
	"inventory/model"
	"inventory/repository"
	"net/http"
)

type purchaseService struct{}

func NewPurchaseService() PurchaseService {
	return new(purchaseService)
}

func (service *purchaseService) InputPurchase(req []model.ReqPurchase) (model.InventoryResponse, error) {
	data, err := repository.NewPurchaseRepository().InputPurchase(req)
	if err != nil {
		return model.InventoryResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
		}, err
	} else {
		return model.InventoryResponse{
			Status:  http.StatusOK,
			Message: "Oke",
			Data:    data,
		}, nil
	}
}

func (service *purchaseService) DetailPurchase(req string) (model.InventoryResponse, error) {
	data, err := repository.NewPurchaseRepository().DetailPurchase(req)
	if err != nil {
		return model.InventoryResponse{
			Status:  http.StatusNotFound,
			Message: "Resource not found",
			Data:    nil,
		}, err
	} else {
		return model.InventoryResponse{
			Status:  http.StatusOK,
			Message: "Oke",
			Data:    data,
		}, nil
	}
}
