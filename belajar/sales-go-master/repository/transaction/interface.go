package transaction

import (
	"sales-go/model"
)

type Repositorier interface {
	GetTransactionByNumber(transactionNumber int) (result []model.TransactionDetail, err error)
	CreateBulkTransactionDetail(voucher model.VoucherRequest, listTransactionDetail []model.TransactionDetail, req model.TransactionDetailBulkRequest) (res []model.TransactionDetail, err error)
}