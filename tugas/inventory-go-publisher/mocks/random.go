package mocks

import "github.com/stretchr/testify/mock"

type RandomMock struct {
	mock.Mock
}

func NewRandom() *RandomMock {
	return &RandomMock{}
}

func (m *RandomMock) Randomizer() string {
	ret := m.Called()
	randomNumber := ret.Get(0).(string)
	return randomNumber
}
