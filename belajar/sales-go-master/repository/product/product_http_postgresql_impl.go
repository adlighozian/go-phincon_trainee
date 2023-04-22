package product

import (
	"context"
	"fmt"
	"sales-go/db"
	"sales-go/model"
	"time"
)

type repositoryhttppostgresql struct {}

func NewPostgreSQLHTTPRepository () *repositoryhttppostgresql {
	return &repositoryhttppostgresql{}
}

func (repo *repositoryhttppostgresql) GetList() (listProduct []model.Product, err error) {
	db := client.NewConnection(client.Database).GetMysqlConnection()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, name, price FROM product`
	stmt, err := db.PrepareContext(ctx, query)
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
	db := client.NewConnection(client.Database).GetMysqlConnection()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, name, price FROM product WHERE name = $1`
	stmt, err := db.PrepareContext(ctx, query)
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
	fmt.Println(client.Database)
	db := client.NewConnection(client.Database).GetMysqlConnection()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	trx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	query := `INSERT INTO product (name, price) VALUES ($1, $2) RETURNING id, name, price`
	stmt, err := db.PrepareContext(ctx, query)
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
