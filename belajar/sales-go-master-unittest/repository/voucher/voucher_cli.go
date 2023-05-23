package voucher

import (
	"errors"
	"sales-go/model"
)

type repositorycli struct {}

func NewCLIRepository() *repositorycli {
	return &repositorycli{}
}

func (repo *repositorycli) GetList() (listVoucher []model.Voucher, err error) {
	return model.VoucherSlice, nil
}

func (repo *repositorycli) GetVoucherByCode(code string) (voucherData model.Voucher, err error) {
	for _, v := range model.VoucherSlice {
		if v.Code == code {
			voucherData = v
		}
	}

	emptyStruct := model.Voucher{}
	if voucherData == emptyStruct {
		err = errors.New("voucher not found")
		return
	}
	return
}

func (repo *repositorycli) Create(req []model.VoucherRequest) (response []model.Voucher, err error) {
	for _, v := range req {
		newData := model.Voucher{
			Id:     len(model.VoucherSlice) + 1,
			Code:   v.Code,
			Persen: v.Persen,
		}
		model.VoucherSlice = append(model.VoucherSlice, newData)
		response = append(response, newData)
	}
	return
}