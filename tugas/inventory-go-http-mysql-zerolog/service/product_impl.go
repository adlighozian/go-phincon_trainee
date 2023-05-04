package service

import (
	"inventory/model"
	"inventory/repository"
	"net/http"
)

type productService struct{}

func NewProductService() ProductService {
	return new(productService)
}

func (service *productService) ShowProduct() (model.InventoryResponse, error) {
	data, err := repository.NewProductRepository().ShowProduct()
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
