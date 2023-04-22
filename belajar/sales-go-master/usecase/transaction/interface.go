package transaction

import (
	"sales-go/model"
)

type TransactionUseCase interface {
	GetTransactionByNumber(number int) (response []model.TransactionDetail, err error)
	CreateBulkTransactionDetail(voucherCode string, req model.TransactionDetailBulkRequest) (response []model.TransactionDetail, err error)
}