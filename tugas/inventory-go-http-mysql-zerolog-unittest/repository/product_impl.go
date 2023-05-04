package repository

import (
	"database/sql"
	"errors"
	"inventory/config/db"
	"inventory/model"
)

type productRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		Db: db,
	}
}

func (repo *productRepository) ShowProduct() ([]model.Product, error) {
	var result []model.Product
	var temp model.Product

	ctx, cancel := db.NewMysqlContext()
	defer cancel()

	query := `SELECT * FROM product`
	rows, err := repo.Db.QueryContext(ctx, query)
	if err != nil {
		return result, errors.New("error melakukan query")
	}

	for rows.Next() {
		rows.Scan(&temp.Id, &temp.Name, &temp.Price, &temp.Stock)
		result = append(result, temp)
	}

	return result, nil
}
