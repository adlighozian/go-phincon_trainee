package transaction

import (
	"errors"
	"math/rand"
	"sales-go/model"
	"time"
)

type repositorycli struct {}

func NewCLIRepository() *repositorycli {
	return &repositorycli{}
}

func (repo *repositorycli) getLastID() (lastID int, err error) {
	if len(model.TransactionDetailSlice) == 0 {
		lastID = 0
	} else {
		lastID = len(model.TransactionDetailSlice)
	}
	return
}

func (repo *repositorycli) GetTransactionByNumber(transactionNumber int) (listSelectedTransaction []model.TransactionDetail, err error) {
	for _, v := range model.TransactionDetailSlice {
		if  v.Transaction.TransactionNumber == transactionNumber {
			listSelectedTransaction = append(listSelectedTransaction, v)
		}
	}
	return
}

func (repo *repositorycli) CreateBulkTransactionDetail(voucher model.VoucherRequest, listTransactionDetail []model.TransactionDetail, req model.TransactionDetailBulkRequest) (res []model.TransactionDetail, err error) {
	lastID, err := repo.getLastID()
	if err != nil {
		return
	}

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
	rand.Seed(time.Now().UnixNano())
	randomInteger := rand.Intn(10000000000)	

	for i, v := range listTransactionDetail {
		newTransaction := model.TransactionDetail{
			Id:       lastID+i+1,
			Item:	  v.Item,
			Price:	  v.Price,
			Quantity: v.Quantity,
			Total:	  v.Total,
			Transaction: model.Transaction{
				Id: 	  		   lastID+1,
				TransactionNumber: randomInteger,
				Name:	  		   req.Name,
				Quantity: 		   quantity,
				Discount: 		   discount,
				Total:	  		   total,
				Pay:	  		   req.Pay,
			},
		}
		model.TransactionDetailSlice = append(model.TransactionDetailSlice, newTransaction)
		res = append(res, newTransaction)
	}

	return
}
