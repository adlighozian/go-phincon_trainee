package service

import (
	"fmt"
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
			Status:  http.StatusBadGateway,
			Message: "Internal Database Error",
			Data:    nil,
		}, err
	} else {
		return model.InventoryResponse{
			Status:  http.StatusOK,
			Message: "oke",
			Data:    data,
		}, nil
	}
}

func (service *purchaseService) DetailPurchase(req string) (model.InventoryResponse, error) {
	fmt.Println("service")
	data, err := repository.NewPurchaseRepository().DetailPurchase(req)
	if err != nil {
		return model.InventoryResponse{
			Status:  http.StatusBadGateway,
			Message: "Internal Database Error",
			Data:    nil,
		}, err
	} else {
		return model.InventoryResponse{
			Status:  http.StatusOK,
			Message: "oke",
			Data:    data,
		}, nil
	}
}
