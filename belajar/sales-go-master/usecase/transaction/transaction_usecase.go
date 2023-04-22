package transaction

import (
	"errors"
	"fmt"
	"sales-go/model"
	"sales-go/repository/product"
	"sales-go/repository/transaction"
	"sales-go/repository/voucher"
)

type usecase struct {
	repo transaction.Repositorier
	productrepo product.Repositorier
	voucherrepo voucher.Repositorier
}

func NewDBHTTPUsecase(
	repository transaction.Repositorier,
	productrepository product.Repositorier,
	voucherrepository voucher.Repositorier,
) *usecase {
	return &usecase{
		repo: repository,
		productrepo: productrepository,
		voucherrepo: voucherrepository,
	}
}

func (uc *usecase) GetTransactionByNumber(number int) (response []model.TransactionDetail, err error) {
	if number < 0 {
		err = errors.New("id must be > 0")
		return
	}

	response, err = uc.repo.GetTransactionByNumber(number)
	if err != nil {
		return
	}
	return
}

func (uc *usecase) CreateBulkTransactionDetail(voucherCode string, req model.TransactionDetailBulkRequest) (response []model.TransactionDetail, err error) {
	listTransactionDetail := []model.TransactionDetail{}
	for _, transaction := range req.Items {
		product, err := uc.productrepo.GetProductByName(transaction.Item)
		if err != nil {
			return []model.TransactionDetail{}, err
		}

		if transaction.Item == "" {
			err = fmt.Errorf("item transaction hould not be empty")
			return []model.TransactionDetail{}, err
		}
	
		if transaction.Quantity < 1 {
			err = fmt.Errorf("quantity transaction should not be negative")
			return []model.TransactionDetail{}, err
		}

		listTransactionDetail = append(listTransactionDetail, model.TransactionDetail{
			Item:     transaction.Item,
			Price:    product.Price,
			Quantity: transaction.Quantity,
			Total:    product.Price * float64(transaction.Quantity),
		})
	}

	var voucher model.VoucherRequest
	if voucherCode == "" {
		voucherData, err := uc.voucherrepo.GetVoucherByCode(voucherCode)
		if err != nil {
			return []model.TransactionDetail{}, err
		} else {
			voucher = model.VoucherRequest{
				Code:   voucherData.Code,
				Persen: voucherData.Persen,
			}
		}
	}
	
	response, err = uc.repo.CreateBulkTransactionDetail(voucher, listTransactionDetail, req)
	if err != nil {
		return
	}
	return
}