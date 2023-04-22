package transaction

import (
	"context"
	"errors"
	"time"

	"sales-go/db"
	"sales-go/helpers/random"
	"sales-go/model"
)

type repositoryhttppostgresql struct {}

func NewPostgreSQLHTTPRepository() *repositoryhttppostgresql {
	return &repositoryhttppostgresql{}
}

func (repo *repositoryhttppostgresql) GetTransactionByNumber(transactionNumber int) (result []model.TransactionDetail, err error) {
	db := client.NewConnection(client.Database).GetMysqlConnection()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	query := `SELECT id, transaction_number, name, quantity, discount, total, pay FROM transaction WHERE transaction_number = $1`
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
	query2 := `SELECT id, item, price, quantity, total FROM transaction_detail WHERE transaction_id = $1`
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

func (repo *repositoryhttppostgresql) CreateBulkTransactionDetail(voucher model.VoucherRequest, listTransactionDetail []model.TransactionDetail, req model.TransactionDetailBulkRequest) (res []model.TransactionDetail, err error) {
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

	query := `INSERT INTO transaction (transaction_number, name, quantity, discount, total, pay) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	var lastIDTransaction int32
	err = stmt.QueryRowContext(ctx, randomInteger, req.Name, quantity, discount, total, req.Pay).Scan(&lastIDTransaction)
	if err != nil {
		trx.Rollback()
		return
	}

	query2 := `INSERT INTO transaction_detail (transaction_id, item, price, quantity, total) values ($1, $2, $3, $4, $5) RETURNING id`
	stmt2, err := db.PrepareContext(ctx, query2)
	if err != nil {
		return
	}

	for _, v := range listTransactionDetail {
		var lastID int32
		err := stmt2.QueryRowContext(ctx, lastIDTransaction, v.Item, v.Price, v.Quantity, v.Total).Scan(&lastID)
		if err != nil {
			trx.Rollback()
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
