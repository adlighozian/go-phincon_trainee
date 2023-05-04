package repository

import (
	"context"
	"errors"
	"fmt"
	"inventory/db"
	"inventory/model"
	"math/rand"
	"time"
)

type purchaseRepository struct{}

func NewPurchaseRepository() PurchaseRepository {
	return new(purchaseRepository)
}

func (repo *purchaseRepository) randomizerPurchase() string {
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

func (repo *purchaseRepository) searchItemPurchase(req string) bool {

	// db := db.GetConnection()
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// query := `SELECT id FROM product WHERE name = $1`
	var isi int = 0
	// rows := db.QueryRowContext(ctx, query, req).Scan(&isi)
	// if rows != nil {
	// 	panic(rows)
	// }

	if isi == 0 {
		return false
	} else {
		return true
	}
}

func (repo *purchaseRepository) InputPurchase(req []model.ReqPurchase) ([]model.PurchaseDetail, error) {
	var dataReturn []model.PurchaseDetail

	db := db.GetConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// INSERT INTO transaction (transaction_number, name, quantity, discount, total, pay) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	fmt.Println("===================debug=======================")
	query := `INSERT INTO purchase (order_number,orang,total) VALUES ($1,$2,$3) RETURNING id`
	query1 := `INSERT INTO purchase_detail (purchase_id,item,price,quantity,total) VALUES ($1,$2,$3,$4,$5) RETURNING id`
	query2 := `INSERT INTO product (name,price,stock) VALUES ($1,$2,$3)`
	query3 := `UPDATE product SET stock = ? WHERE name = ?;`

	txr, errs := db.BeginTx(ctx, nil)
	stmt, _ := txr.PrepareContext(ctx, query)
	stmt1, _ := txr.PrepareContext(ctx, query1)
	stmt2, _ := txr.PrepareContext(ctx, query2)
	stmt3, _ := txr.PrepareContext(ctx, query3)

	for _, v := range req {
		fmt.Println("===================debug2=======================")
		order := repo.randomizerPurchase()
		if !repo.searchItemPurchase(v.Item) {
			// bikin baru
			fmt.Println("===================debug4=======================")
			var purchase_id int32
			purchase := stmt.QueryRowContext(ctx, order, v.From, v.Total).Scan(&purchase_id)
			if purchase != nil {
				panic(purchase)
			}
			fmt.Println("===================debug5=======================", purchase_id)
			var purchaseDetail_id int32
			purchaseDetail := stmt1.QueryRowContext(ctx, purchase_id, v.Item, v.Price, v.Total, v.Total).Scan(&purchaseDetail_id)
			if purchaseDetail != nil {
				panic(purchaseDetail)
			}

			stmt2.ExecContext(ctx, v.Item, v.Price, v.Total)

			orderDetail := model.PurchaseDetail{
				Id:       int(purchaseDetail_id),
				Item:     v.Item,
				Price:    v.Price,
				Quantity: v.Total,
				Total:    v.Total,
				Purchase: model.Purchase{
					Id:          int(purchase_id),
					OrderNumber: order,
					From:        v.From,
					Total:       v.Total,
				},
			}
			dataReturn = append(dataReturn, orderDetail)

		} else {
			fmt.Println("===================debug3=======================")
			// update barang
			var stock int
			query := `SELECT stock FROM product WHERE name = ? `
			rows, _ := db.QueryContext(ctx, query, v.Item)
			for rows.Next() {
				rows.Scan(&stock)
			}
			Titem := v.Total + stock
			stmt3.ExecContext(ctx, Titem, v.Item)

			result, _ := stmt.ExecContext(ctx, order, v.From, v.Total)
			lastInsertId, err := result.LastInsertId()
			if err != nil {
				panic(err)
			}

			results, _ := stmt1.ExecContext(ctx, lastInsertId, v.Item, v.Price, v.Total, Titem)
			lastInsertIds, err := results.LastInsertId()
			if err != nil {
				panic(err)
			}

			orderDetails := model.PurchaseDetail{
				Id:       int(lastInsertIds),
				Item:     v.Item,
				Price:    v.Price,
				Quantity: v.Total,
				Total:    Titem,
				Purchase: model.Purchase{
					Id:          int(lastInsertId),
					OrderNumber: order,
					From:        v.From,
					Total:       v.Total,
				},
			}
			dataReturn = append(dataReturn, orderDetails)
		}
	}

	if errs != nil {
		txr.Rollback()
	} else {
		txr.Commit()
	}

	if dataReturn == nil {
		return nil, errors.New("gagal menambahkan data")
	} else {
		return dataReturn, nil
	}
}

func (repo *purchaseRepository) DetailPurchase(req string) (model.PurchaseDetail, error) {
	db := db.GetConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	var kotak model.PurchaseDetail

	modelp := new(model.Purchase)
	query := `SELECT id, order_number, orang, total FROM purchase WHERE order_number = ? `
	rows := db.QueryRowContext(ctx, query, req)
	rows.Scan(&modelp.Id, &modelp.OrderNumber, &modelp.From, &modelp.Total)

	modelpd := new(model.PurchaseDetail)
	sqlQuery := "SELECT id, item, price, quantity, total FROM purchase_detail WHERE purchase_id = ?"
	row := db.QueryRowContext(ctx, sqlQuery, modelp.Id)
	row.Scan(&modelpd.Id, &modelpd.Item, &modelpd.Price, &modelpd.Quantity, &modelpd.Total)

	kotak = model.PurchaseDetail{
		Id:       modelpd.Id,
		Item:     modelpd.Item,
		Price:    modelpd.Price,
		Quantity: modelpd.Quantity,
		Total:    modelpd.Total,
		Purchase: model.Purchase{
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
