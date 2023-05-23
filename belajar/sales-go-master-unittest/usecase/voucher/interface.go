package voucher

import (
	"sales-go/model"
)

type VoucherUseCase interface {
	GetList() (response []model.Voucher, err error)
	GetVoucherByCode(name string) (response model.Voucher, err error)
	Create(req []model.VoucherRequest) (response []model.Voucher, err error)
}