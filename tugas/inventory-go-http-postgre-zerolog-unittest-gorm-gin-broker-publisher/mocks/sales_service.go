package mocks

import (
	"inventory/model"

	"github.com/stretchr/testify/mock"
)

type servSalesMock struct {
	mock.Mock
}

func NewServSalesMock() *servSalesMock {
	return &servSalesMock{}
}

func (m *servSalesMock) InputSales(req []model.ReqSales) ([]model.SalesDetail, error) {
	ret := m.Called(req)
	res := ret.Get(0).([]model.SalesDetail)
	err := ret.Error(1)
	return res, err
}

func (m *servSalesMock) DetailSales(req string) (model.SalesDetail, error) {
	ret := m.Called(req)
	res := ret.Get(0).(model.SalesDetail)
	err := ret.Error(1)
	return res, err
}
