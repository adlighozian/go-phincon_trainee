// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	model "contact-go/model"

	mock "github.com/stretchr/testify/mock"
)

// ContactUsecase is an autogenerated mock type for the ContactUsecase type
type ContactUsecase struct {
	mock.Mock
}

// Add provides a mock function with given fields: req
func (_m *ContactUsecase) Add(req *model.ContactRequest) (*model.Contact, error) {
	ret := _m.Called(req)

	var r0 *model.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.ContactRequest) (*model.Contact, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*model.ContactRequest) *model.Contact); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.ContactRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *ContactUsecase) Delete(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Detail provides a mock function with given fields: id
func (_m *ContactUsecase) Detail(id int64) (*model.Contact, error) {
	ret := _m.Called(id)

	var r0 *model.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*model.Contact, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) *model.Contact); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *ContactUsecase) List() ([]model.Contact, error) {
	ret := _m.Called()

	var r0 []model.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.Contact, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Contact); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, req
func (_m *ContactUsecase) Update(id int64, req *model.ContactRequest) (*model.Contact, error) {
	ret := _m.Called(id, req)

	var r0 *model.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, *model.ContactRequest) (*model.Contact, error)); ok {
		return rf(id, req)
	}
	if rf, ok := ret.Get(0).(func(int64, *model.ContactRequest) *model.Contact); ok {
		r0 = rf(id, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, *model.ContactRequest) error); ok {
		r1 = rf(id, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewContactUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewContactUsecase creates a new instance of ContactUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewContactUsecase(t mockConstructorTestingTNewContactUsecase) *ContactUsecase {
	mock := &ContactUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}