package mocks

import (
	"github.com/stretchr/testify/mock"
)

type ControllerMock struct {
	mock.Mock
}

func NewControllerMock() *ControllerMock {
	return &ControllerMock{}
}

func (m *ControllerMock) ComparePassword(hashedPassword string, password string) (err error) {
	ret := m.Called(hashedPassword, password)
	err = ret.Error(0)
	return err
}
