package mocks

import "github.com/stretchr/testify/mock"

type PublisherMock struct {
	mock.Mock
}

func NewPublisher() *PublisherMock {
	return &PublisherMock{}
}

func (m *PublisherMock) PubPurchase(body interface{}) (err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(body)
	err = ret.Error(0)
	return
}
