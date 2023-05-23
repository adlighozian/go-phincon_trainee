package service

import (
	"inventory/model"
	"inventory/repository/purchase"
	"log"
	"net/http"
)

type purchaseService struct {
	repository purchase.PurchaseRepository
}

func NewPurchaseService(repo purchase.PurchaseRepository) PurchaseService {
	return &purchaseService{
		repository: repo,
	}
}

func (s *purchaseService) InputPurchase(req []model.ReqPurchase) (model.InventoryResponse, error) {
	log.Println("purchase service")
	data, err := s.repository.InputPurchase(req)
	if err != nil {
		return model.InventoryResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
		}, err
	} else {
		return model.InventoryResponse{
			Status:  http.StatusOK,
			Message: "OK",
			Data:    data,
		}, nil
	}
}

func (s *purchaseService) DetailPurchase(req string) (model.InventoryResponse, error) {
	log.Println("purchase service")
	data, err := s.repository.DetailPurchase(req)
	if err != nil {
		return model.InventoryResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
		}, err
	} else {
		return model.InventoryResponse{
			Status:  http.StatusOK,
			Message: "OK",
			Data:    data,
		}, nil
	}
}
