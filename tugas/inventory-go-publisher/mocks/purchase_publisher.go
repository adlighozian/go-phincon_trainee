package mocks

import "github.com/stretchr/testify/mock"

type PurchaseMock struct {
	mock.Mock
}

func NewPublisher() *PurchaseMock {
	return &PurchaseMock{}
}

func (m *PurchaseMock) PubPurchase(body interface{}) (err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(body)
	err = ret.Error(0)
	return
}
