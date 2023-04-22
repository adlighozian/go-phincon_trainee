package product

import (
	"sales-go/model"
)

type ProductUseCase interface {
	GetList() (response []model.Product, err error)
	GetProductByName(name string) (response model.Product, err error)
	Create(req []model.ProductRequest) (response []model.Product, err error)
}