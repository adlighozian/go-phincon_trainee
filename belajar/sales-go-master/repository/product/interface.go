package product

import (
	"sales-go/model"
)

type Repositorier interface {
	GetList() (listProduct []model.Product, err error)
	GetProductByName(name string) (productData model.Product, err error)
	Create(req []model.ProductRequest) (result []model.Product, err error)
}
