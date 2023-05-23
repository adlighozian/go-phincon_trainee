package random

import (
	"github.com/stretchr/testify/mock"
)

type RandomMock struct {
	mock.Mock
}

func NewRandom() *RandomMock {
	return &RandomMock{}
}

func (m *RandomMock) RandomString(length int) (int, error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(length)
	randomNumber := ret.Get(0).(int)
	err := ret.Error(1)
	return randomNumber, err
}
