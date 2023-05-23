package product

import (
	"inventory/model"
)

type ProductRepository interface {
	ShowProduct() ([]model.Product, error)
}
