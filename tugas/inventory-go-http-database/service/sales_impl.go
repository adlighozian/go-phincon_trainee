package service

import (
	"inventory/model"
	"inventory/repository"
	"net/http"
)

type salesService struct{}

func NewSalesService() SalesService {
	return new(salesService)
}

func (service *salesService) InputSales(req []model.ReqSales) (model.InventoryResponse, error) {
	// fmt.Println("service : sales input")
	data, err := repository.NewSalesRepository().InputSales(req)
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

func (service *salesService) ShowSales(req string) (model.InventoryResponse, error) {
	data, err := repository.NewSalesRepository().ShowSales(req)
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
