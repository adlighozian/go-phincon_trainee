package service

import (
	"inventory/model"
	"inventory/repository/product"
	"net/http"
)

type productService struct {
	repository product.ProductRepository
}

func NewProductService(repo product.ProductRepository) ProductService {
	return &productService{
		repository: repo,
	}
}

func (service *productService) ShowProduct() (model.InventoryResponse, error) {
	data, err := service.repository.ShowProduct()
	if err != nil {
		return model.InventoryResponse{
			Status:  http.StatusBadGateway,
			Message: "Bad Gateway",
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
