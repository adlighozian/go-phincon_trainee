package mocks

import (
	"inventory/model"

	"github.com/stretchr/testify/mock"
)

type SalesControllerMock struct {
	mock.Mock
}

func NewSalesControllerMock() *SalesControllerMock {
	return &SalesControllerMock{}
}

func (m *SalesControllerMock) InputSales(req []model.ReqSales) (model.InventoryResponse, error) {
	ret := m.Called(req)
	res := ret.Get(0).(model.InventoryResponse)
	err := ret.Error(1)
	return res, err
}

func (m *SalesControllerMock) DetailSales(req string) (model.InventoryResponse, error) {
	ret := m.Called(req)
	res := ret.Get(0).(model.InventoryResponse)
	err := ret.Error(1)
	return res, err
}
