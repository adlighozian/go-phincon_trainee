package mocks

import (
	"contact-go/model"

	"github.com/stretchr/testify/mock"
)

type HandlerMock struct {
	mock.Mock
}

func NewHandlerMock() *HandlerMock {
	return &HandlerMock{}
}

//mock usecase

func (m *HandlerMock) List() (model.ContactResponse, error) {
	ret := m.Called()
	response := ret.Get(0).(model.ContactResponse)
	err := ret.Error(1)

	return response, err

}

func (m *HandlerMock) Add(req []model.ContactRequest) (model.ContactResponse, error) {
	ret := m.Called(req)
	response := ret.Get(0).(model.ContactResponse)
	err := ret.Error(1)

	return response, err
}

func (m *HandlerMock) Update(id int, req model.ContactRequest) (model.ContactResponse, error) {
	ret := m.Called(id, req)
	response := ret.Get(0).(model.ContactResponse)
	err := ret.Error(1)

	return response, err
}

func (m *HandlerMock) Delete(id int) (model.ContactResponse, error) {
	ret := m.Called(id)
	response := ret.Get(0).(model.ContactResponse)
	err := ret.Error(1)

	return response, err
}
