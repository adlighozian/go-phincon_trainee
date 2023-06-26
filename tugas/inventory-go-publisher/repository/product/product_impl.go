package product

import (
	"errors"
	"inventory/model"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(dbs *gorm.DB) ProductRepository {
	return &productRepository{
		db: dbs,
	}
}

func (repo *productRepository) ShowProduct() ([]model.Product, error) {
	var productrepo []model.Product

	query := "SELECT * FROM products"

	repo.db.Raw(query).Scan(&productrepo)

	if productrepo == nil {
		return productrepo, errors.New("")
	}

	return productrepo, nil

}
