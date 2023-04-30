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

func (service *salesService) DetailSales(req string) (model.InventoryResponse, error) {
	data, err := repository.NewSalesRepository().ShowSales(req)
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
