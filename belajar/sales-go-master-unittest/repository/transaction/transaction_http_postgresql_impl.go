package transaction

import (
	"context"
	"database/sql"
	"fmt"
	"errors"
	"time"

	"sales-go/helpers/random"
	"sales-go/model"
	"sales-go/publisher"
)

type repositoryhttppostgresql struct {	
	db 			*sql.DB
	publisher	publisher.PublisherInterface
	random		random.RandomInterface
}

func NewPostgreSQLHTTPRepository(db *sql.DB, publisher publisher.PublisherInterface, random random.RandomInterface) *repositoryhttppostgresql {
	return &repositoryhttppostgresql{
		db: 		db,
		publisher:	publisher,
		random: 	random,
	}
}

func (repo *repositoryhttppostgresql) GetTransactionByNumber(transactionNumber int) (result []model.TransactionDetail, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// GetTransaction
	query := `SELECT id, transaction_number, name, quantity, discount, total, pay FROM transaction WHERE transaction_number = $1`
	stmt, err := repo.db.PrepareContext(ctx, query)
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
	stmt2, err := repo.db.PrepareContext(ctx, query2)
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
	// generate random integer
	randomInteger, err := repo.random.RandomString(9)
	if err != nil {
		return
	}
	fmt.Println("RANDOM : ", randomInteger)

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
	
	data := model.TransactionRabbitMQData{
		Name:					req.Name,
		RandomInteger:			randomInteger,
		Quantity:				quantity,
		Total:					total,
		Discount:				discount,
		Pay:					req.Pay,
		ListTransactionDetail:	listTransactionDetail,
	}

	// publish data to RabbitMQ
	err = repo.publisher.Publish(data)
	if err != nil {
		err = fmt.Errorf("error publish data to RabbitMQ : %s", err.Error())
		return
	}

	for _, v := range data.ListTransactionDetail {
		res = append(res,  model.TransactionDetail{
			Item:	  v.Item,
			Price:	  v.Price,
			Quantity: v.Quantity,
			Total:	  v.Total,
			Transaction: model.Transaction{
				TransactionNumber: data.RandomInteger,
				Name:	  		   data.Name,
				Quantity: 		   data.Quantity,
				Discount: 		   data.Discount,
				Total:	  		   data.Total,
				Pay:	  		   data.Pay,
			},
		})
	}

	return
}
