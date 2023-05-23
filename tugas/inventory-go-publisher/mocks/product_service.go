package mocks

import (
	"inventory/model"

	"github.com/stretchr/testify/mock"
)

type ServProductMock struct {
	mock.Mock
}

func NewServProductMock() *ServProductMock {
	return &ServProductMock{}
}

func (m *ServProductMock) ShowProduct() ([]model.Product, error) {
	ret := m.Called()
	result := ret.Get(0).([]model.Product)
	err := ret.Error(1)
	return result, err
}
