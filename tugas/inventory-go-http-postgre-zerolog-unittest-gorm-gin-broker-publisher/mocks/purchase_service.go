package mocks

import (
	"inventory/model"

	"github.com/stretchr/testify/mock"
)

type ServPurchaseMock struct {
	mock.Mock
}

func NewServPurchaseMock() *ServPurchaseMock {
	return &ServPurchaseMock{}
}

func (m *ServPurchaseMock) InputPurchase(req []model.ReqPurchase) ([]model.PurchaseDetail, error) {
	ret := m.Called(req)
	result := ret.Get(0).([]model.PurchaseDetail)
	err := ret.Error(1)
	return result, err
}

func (m *ServPurchaseMock) DetailPurchase(req string) (model.PurchaseDetail, error) {
	ret := m.Called(req)
	result := ret.Get(0).(model.PurchaseDetail)
	err := ret.Error(1)
	return result, err
}
