package voucher

import (
	"sales-go/model"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct{
	mock.Mock
}

func NewVoucherRepoMock() *RepoMock {
	return &RepoMock{}
}


func (m *RepoMock) GetList() (listVoucher []model.Voucher, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called()
	listVoucher = ret.Get(0).([]model.Voucher)
	err = ret.Error(1)
	return listVoucher, err
}

func (m *RepoMock) GetVoucherByCode(code string) (voucherData model.Voucher, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(code)
	voucherData = ret.Get(0).(model.Voucher)
	err = ret.Error(1)
	return voucherData, err
}

func (m *RepoMock) Create(req []model.VoucherRequest) (response []model.Voucher, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(req)
	response = ret.Get(0).([]model.Voucher)
	err = ret.Error(1)
	return response, err
}
