package transaction

import (
	"context"
	"errors"
	"fmt"
	"time"

	"sales-go/db"
	"sales-go/helpers/random"
	"sales-go/model"
)

type repositoryhttpmysql struct {}

func NewMySQLHTTPRepository() *repositoryhttpmysql {
	return &repositoryhttpmysql{}
}

func (repo *repositoryhttpmysql) GetTransactionByNumber(transactionNumber int) (result []model.TransactionDetail, err error) {
	db := client.NewConnection(client.Database).GetMysqlConnection()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	query := `SELECT id, transaction_number, name, quantity, discount, total, pay FROM transaction WHERE transaction_number = ?`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.QueryContext(ctx, transactionNumber)
	if err != nil {
		return
	}

	var transaction model.Transaction
	for res.Next() {
		res.Scan(&transaction.Id, &transaction.TransactionNumber, &transaction.Name, &transaction.Quantity, &transaction.Discount, &transaction.Total, &transaction.Pay)
	}

	// GetTransactionDetail
	query2 := `SELECT id, item, price, quantity, total FROM transaction_detail WHERE transaction_id = ?`
	stmt2, err := db.PrepareContext(ctx, query2)
	if err != nil {
		return
	}

	res2, err := stmt2.QueryContext(ctx, transaction.Id)
	if err != nil {
		return
	}

	for res2.Next() {
		var temp model.TransactionDetail
		res2.Scan(&temp.Id, &temp.Item, &temp.Price, &temp.Quantity, &temp.Total)
		// append transaction in each of transaction detail
		temp.Transaction = transaction
		result = append(result, temp)
	}

	return
}

func (repo *repositoryhttpmysql) CreateBulkTransactionDetail(voucher model.VoucherRequest, listTransactionDetail []model.TransactionDetail, req model.TransactionDetailBulkRequest) (res []model.TransactionDetail, err error) {
	// sum all quantity and total
	var quantity int
	var total float64
	for _, item := range listTransactionDetail {
		quantity = quantity + item.Quantity
		total = total + item.Total
	}

	// discount calculation
	var discount float64
	if total > 300000 && voucher.Persen > 0 {
		discount = voucher.Persen/100
		total = total*(1-discount)
	}

	if req.Pay < total {
		err = errors.New("pay must be > total")
		return
	}

	// generate random integer
	randomInteger, err := random.RandomString(9)
    if err != nil {
        return
    }
	
	db := client.NewConnection(client.Database).GetMysqlConnection()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	trx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	fmt.Println("PASS 1")
	query := `INSERT INTO transaction (transaction_number, name, quantity, discount, total, pay) values (?, ?, ?, ?, ?, ?)`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	resInsert, err := stmt.ExecContext(ctx, randomInteger, req.Name, quantity, discount, total, req.Pay)
	if err != nil {
		trx.Rollback()
		return
	}

	lastIDTransaction, err := resInsert.LastInsertId()
	if err != nil {
		return
	}

	query2 := `INSERT INTO transaction_detail (transaction_id, item, price, quantity, total) values (?, ?, ?, ?, ?)`
	stmt2, err := db.PrepareContext(ctx, query2)
	if err != nil {
		return
	}

	for _, v := range listTransactionDetail {
		res2, err := stmt2.ExecContext(ctx, lastIDTransaction, v.Item, v.Price, v.Quantity, v.Total)
		if err != nil {
			trx.Rollback()
			return []model.TransactionDetail{}, err
		}

		lastID, err := res2.LastInsertId()
		if err != nil {
			return []model.TransactionDetail{}, err
		}

		newTransaction := model.TransactionDetail{
			Id:       int(lastID),
			Item:	  v.Item,
			Price:	  v.Price,
			Quantity: v.Quantity,
			Total:	  v.Total,
			Transaction: model.Transaction{
				Id: 	  		   int(lastIDTransaction),
				TransactionNumber: randomInteger,
				Name:	  		   req.Name,
				Quantity: 		   quantity,
				Discount: 		   discount,
				Total:	  		   total,
				Pay:	  		   req.Pay,
			},
		}
		res = append(res, newTransaction)
	}

	trx.Commit()

	return
}
