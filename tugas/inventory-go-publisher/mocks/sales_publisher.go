package mocks

import "github.com/stretchr/testify/mock"

type SalesMock struct {
	mock.Mock
}

func NewSales() *SalesMock {
	return &SalesMock{}
}

func (m *SalesMock) PubSales(body interface{}) (err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(body)
	err = ret.Error(0)
	return
}
