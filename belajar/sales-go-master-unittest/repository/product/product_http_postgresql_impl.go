package product

import (
	"context"
	"database/sql"
	"sales-go/model"
	"time"
)

type repositoryhttppostgresql struct {
	db *sql.DB
}

func NewPostgreSQLHTTPRepository(db *sql.DB) *repositoryhttppostgresql {
	return &repositoryhttppostgresql{
		db: db,
	}
}

func (repo *repositoryhttppostgresql) GetList() (listProduct []model.Product, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, name, price FROM product`
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.QueryContext(ctx)
	if err != nil {
		return
	}

	for res.Next() {
		var temp model.Product
		res.Scan(&temp.Id, &temp.Name, &temp.Price)

		listProduct = append(listProduct, temp)
	}

	return
}

func (repo *repositoryhttppostgresql) GetProductByName(name string) (productData model.Product, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, name, price FROM product WHERE name = $1`
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.QueryContext(ctx, name)
	if err != nil {
		return
	}

	for res.Next() {
		res.Scan(&productData.Id, &productData.Name, &productData.Price)
	}
	return
}

func (repo *repositoryhttppostgresql) Create(req []model.ProductRequest) (result []model.Product, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	trx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	query := `INSERT INTO product (name, price) VALUES ($1, $2) RETURNING id, name, price`
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	
	for _, v := range req {
		var temp model.Product
		err = stmt.QueryRowContext(ctx, v.Name, v.Price).Scan(&temp.Id, &temp.Name, &temp.Price)
		if err != nil {
			trx.Rollback()
			return []model.Product{}, err
		}
		result = append(result, temp)
	}

	trx.Commit()

	return
}
