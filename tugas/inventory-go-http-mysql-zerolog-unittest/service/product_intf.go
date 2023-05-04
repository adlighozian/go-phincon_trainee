package service

import "inventory/model"

type ProductService interface {
	ShowProduct() (model.InventoryResponse, error)
}
