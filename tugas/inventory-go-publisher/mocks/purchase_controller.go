package mocks

import (
	"inventory/model"

	"github.com/stretchr/testify/mock"
)

type PurchaseControllerMock struct {
	mock.Mock
}

func NewPurchaseControllerMock() *PurchaseControllerMock {
	return &PurchaseControllerMock{}
}

func (m *PurchaseControllerMock) InputPurchase(req []model.ReqPurchase) (model.InventoryResponse, error) {
	ret := m.Called(req)
	res := ret.Get(0).(model.InventoryResponse)
	err := ret.Error(1)
	return res, err
}

func (m *PurchaseControllerMock) DetailPurchase(req string) (model.InventoryResponse, error) {
	ret := m.Called(req)
	res := ret.Get(0).(model.InventoryResponse)
	err := ret.Error(1)
	return res, err
}
