package repository

import (
	"inventory/model"
)

type ProductRepository interface {
	ShowProduct() []model.Product
	DecodeProduct() []model.Product
	EncodeProduct()
}
