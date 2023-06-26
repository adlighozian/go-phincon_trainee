package mocks

import (
	"inventory/model"

	"github.com/stretchr/testify/mock"
)

type ProductControllerMock struct {
	mock.Mock
}

func NewProductControllerMock() *ProductControllerMock {
	return &ProductControllerMock{}
}

func (m ProductControllerMock) ShowProduct() (model.InventoryResponse, error) {
	ret := m.Called()
	res := ret.Get(0).(model.InventoryResponse)
	err := ret.Error(1)
	return res, err
}
