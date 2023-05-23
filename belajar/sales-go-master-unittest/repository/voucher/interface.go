package voucher

import (
	"sales-go/model"
)


type Repositorier interface {
	GetList() (listVoucher []model.Voucher, err error)
	GetVoucherByCode(code string) (voucherData model.Voucher, err error)
	Create(req []model.VoucherRequest) (response []model.Voucher, err error)
}