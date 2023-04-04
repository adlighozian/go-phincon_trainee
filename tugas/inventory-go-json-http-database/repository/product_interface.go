package repository

import (
	"inventory/model"
)

type ProductRepository interface {
	ShowProduct() []model.Product
	DecodeProduct() []model.Product
	EncodeProduct()
	GetIdProduct() int
	SearchItem(param string) bool
}
