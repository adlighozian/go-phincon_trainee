package repository

import (
	"inventory/model"
)

type productRepository struct {
}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (repo *productRepository) ShowProduct() []model.Product {
	return model.Products
}
