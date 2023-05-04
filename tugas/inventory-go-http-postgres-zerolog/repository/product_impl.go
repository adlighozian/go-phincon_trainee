package repository

import (
	"context"
	"errors"
	"fmt"
	"inventory/db"
	"inventory/model"
	"time"
)

type productRepository struct {
}

func NewProductRepository() ProductRepository {
	return new(productRepository)
}

func (repo *productRepository) ShowProduct() ([]model.Product, error) {
	var result []model.Product
	var temp model.Product

	db := db.GetConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, name, price, stock FROM product`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("error query", err)
		return result, errors.New("error melakukan query")
	}

	for rows.Next() {
		rows.Scan(&temp.Id, &temp.Name, &temp.Price, &temp.Stock)
		result = append(result, temp)
	}

	return result, nil
}
