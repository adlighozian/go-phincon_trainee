package service

import (
	"inventory/model"
	"inventory/repository/sales"
	"net/http"
)

type salesService struct {
	repository sales.SalesRepository
}

func NewSalesService(repo sales.SalesRepository) SalesService {
	return &salesService{
		repository: repo,
	}
}

func (s *salesService) InputSales(req []model.ReqSales) (model.InventoryResponse, error) {
	data, err := s.repository.InputSales(req)
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

func (s *salesService) DetailSales(req string) (model.InventoryResponse, error) {
	data, err := s.repository.DetailSales(req)
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
