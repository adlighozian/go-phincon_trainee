package transaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sales-go/model"
	"time"
)

type repository struct {}

func NewJsonRepository() *repository {
	return &repository{}
}

func (repo *repository) getLastID() (lastID int, err error) {
	listTransaction, err := repo.GetListTransactionDetail()
	if err != nil {
		return
	}

	if len(listTransaction) == 0 {
		lastID = 0
	} else {
		lastID = len(listTransaction)
	}
	return
}

func (repo *repository) GetListTransactionDetail() (res []model.TransactionDetail, err error) {
	reader, err := os.Open("data/transaction.json")
	if err != nil {
		err = errors.New(fmt.Sprintf("[ERROR] os open transaction detail json %s", err.Error()))
		return
	}

	decoder := json.NewDecoder(reader)
	decoder.Decode(&res)

	return
}

func (repo *repository) UpdateJSONTransactionDetail(req []model.TransactionDetail) (err error) {
	writerJson, err := os.Create("data/transaction.json")
	if err != nil {
		err = errors.New(fmt.Sprintf("[ERROR] os create transaction json %s", err.Error()))
		return
	}
	encodeToJson := json.NewEncoder(writerJson)
	encodeToJson.Encode(req)

	writerTxt, err := os.Create("data/transaction.txt")
	if err != nil { 
		err = errors.New(fmt.Sprintf("[ERROR] os create transaction txt %s", err.Error()))
		return
	}
	encodeToTxt := json.NewEncoder(writerTxt)
	encodeToTxt.Encode(req)

	return
}

func (repo *repository) GetTransactionByNumber(transactionNumber int) (listSelectedTransaction []model.TransactionDetail, err error) {
	listTransaction, err := repo.GetListTransactionDetail()
	if err != nil {
		return
	}

	for _, v := range listTransaction {
		if  v.Transaction.TransactionNumber == transactionNumber {
			listSelectedTransaction = append(listSelectedTransaction, v)
		}
	}
	return
}

func (repo *repository) CreateBulkTransactionDetail(voucher model.VoucherRequest, listTransactionDetail []model.TransactionDetail, req model.TransactionDetailBulkRequest) (res []model.TransactionDetail, err error) {
	listTransaction, err := repo.GetListTransactionDetail()
	if err != nil {
		return
	}

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
		listTransaction = append(listTransaction, newTransaction)
		res = append(res, newTransaction)

		err = repo.UpdateJSONTransactionDetail(listTransaction)
		if err != nil {
			return
		}
	}

	return
}
