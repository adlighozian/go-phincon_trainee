package voucher

import (
	"errors"
	"fmt"
	"sales-go/model"
	"sales-go/repository/voucher"
)

type usecase struct {
	repo voucher.Repositorier
}

func NewDBHTTPUsecase(repository voucher.Repositorier) *usecase {
	return &usecase{
		repo: repository,
	}
}

func (uc *usecase) GetList() (response []model.Voucher, err error) {
	return uc.repo.GetList()
}

func (uc *usecase) GetVoucherByCode(name string) (response model.Voucher, err error) {
	response, err = uc.repo.GetVoucherByCode(name)
	if err != nil {
		return
	}

	emptyStruct := model.Voucher{}
	if response == emptyStruct{
		err = errors.New("voucher not found")
		return
	}
	return
}

func (uc *usecase) Create(req []model.VoucherRequest) (response []model.Voucher, err error) {
	for _, voucher := range req {
		if voucher.Code == "" {
			err = fmt.Errorf("voucher %s : code should not be empty", voucher.Code)
			return
		} else if voucher.Persen <= 0 {
			err = fmt.Errorf("voucher %s : persen should be > 0", voucher.Code)
			return
		} else {
			_, err = uc.GetVoucherByCode(voucher.Code)
			if err != nil {
				continue
			} else {
				fmt.Println(err)
				err = fmt.Errorf("voucher %s already exist", voucher.Code)
				return
			}
		}
	}

	response, err = uc.repo.Create(req)
	if err != nil {
		return
	}
	return
}