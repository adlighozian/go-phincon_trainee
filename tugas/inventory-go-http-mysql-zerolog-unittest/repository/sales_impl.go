package repository

import (
	"context"
	"errors"
	"inventory/db"
	"inventory/model"
	"math/rand"
	"time"
)

type salesRepository struct {
}

func NewSalesRepository() SalesRepository {
	return new(salesRepository)
}

func (repo *salesRepository) randomizerSales() string {
	time.Sleep(1 * time.Second)
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 7)

	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	rand := string(b)
	return rand
}

func (repo *salesRepository) searchItemSales(req string) bool {

	db := db.GetConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id FROM product WHERE name = ? `
	rows, err := db.QueryContext(ctx, query, req)
	if err != nil {
		panic(err)
	}

	var isi int = 0
	for rows.Next() {
		isi++
	}

	if isi == 0 {
		return false
	} else {
		return true
	}
}

func (repo *salesRepository) InputSales(req []model.ReqSales) ([]model.SalesDetail, error) {
	// fmt.Println("repository : sales input")
	var dataReturn []model.SalesDetail

	db := db.GetConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	updateSales := `UPDATE product SET stock = ? WHERE name = ?;`
	insertSalesDetail := `INSERT INTO sales_detail(sales_id,item,price,quantity,total) VALUES (?,?,?,?,?)`
	insertSales := `INSERT INTO sales(order_number,orang,total) VALUES (?,?,?)`
	selectStockProduct := `SELECT stock FROM product WHERE name = ? `

	txr, errs := db.BeginTx(ctx, nil)
	stmtUpdate, _ := txr.PrepareContext(ctx, updateSales)
	stmtSalesDetail, _ := txr.PrepareContext(ctx, insertSalesDetail)
	stmtSales, _ := txr.PrepareContext(ctx, insertSales)
	stmtStockProduct, _ := txr.PrepareContext(ctx, selectStockProduct)

	for _, v := range req {
		order := repo.randomizerSales()
		if !repo.searchItemSales(v.Item) {
			continue
		}

		var stock int

		rows, _ := stmtStockProduct.QueryContext(ctx, v.Item)
		for rows.Next() {
			rows.Scan(&stock)
		}

		if stock < v.Total {
			continue
		}
		Titem := stock - v.Total
		stmtUpdate.ExecContext(ctx, Titem, v.Item)

		result, err := stmtSales.ExecContext(ctx, order, v.From, v.Total)
		if err != nil {
			panic(err)
		}
		lastInsertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		results, err := stmtSalesDetail.ExecContext(ctx, lastInsertId, v.Item, v.Price, v.Total, Titem)
		if err != nil {
			panic(err)
		}
		lastInsertIds, err := results.LastInsertId()
		if err != nil {
			panic(err)
		}

		salesDetails := model.SalesDetail{
			Id:       int(lastInsertIds),
			Item:     v.Item,
			Price:    v.Price,
			Quantity: v.Total,
			Total:    Titem,
			Sales: model.Sales{
				Id:          int(lastInsertId),
				OrderNumber: order,
				From:        v.From,
				Total:       v.Total,
			},
		}

		dataReturn = append(dataReturn, salesDetails)
	}

	if errs != nil {
		txr.Rollback()
	} else {
		txr.Commit()
	}

	if dataReturn == nil {
		return dataReturn, errors.New("data tidak ditemukan")
	} else {
		return dataReturn, nil
	}

}

func (repo *salesRepository) ShowSales(req string) (model.SalesDetail, error) {
	db := db.GetConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	var kotak model.SalesDetail

	modelp := new(model.Sales)
	query := `SELECT id, order_number, orang, total FROM sales WHERE order_number = ? `
	rows := db.QueryRowContext(ctx, query, req)
	rows.Scan(&modelp.Id, &modelp.OrderNumber, &modelp.From, &modelp.Total)

	modelpd := new(model.SalesDetail)
	sqlQuery := "SELECT id, item, price, quantity, total FROM sales_detail WHERE sales_id = ?"
	row := db.QueryRowContext(ctx, sqlQuery, modelp.Id)
	row.Scan(&modelpd.Id, &modelpd.Item, &modelpd.Price, &modelpd.Quantity, &modelpd.Total)

	kotak = model.SalesDetail{
		Id:       modelpd.Id,
		Item:     modelpd.Item,
		Price:    modelpd.Price,
		Quantity: modelpd.Quantity,
		Total:    modelpd.Total,
		Sales: model.Sales{
			Id:          modelp.Id,
			OrderNumber: modelp.OrderNumber,
			From:        modelp.From,
			Total:       modelp.Total,
		},
	}

	if modelp.Id == 0 {
		return kotak, errors.New("gagal mencari")
	} else {
		return kotak, nil
	}

}
