package mocks

import (
	model "contact-go/model"

	mock "github.com/stretchr/testify/mock"
)

// ContactRepository is an autogenerated mock type for the ContactRepository type
type ContactRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: contact
func (_m *ContactRepository) Add(contact *model.Contact) (*model.Contact, error) {
	ret := _m.Called(contact)

	var r0 *model.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Contact) (*model.Contact, error)); ok {
		return rf(contact)
	}
	if rf, ok := ret.Get(0).(func(*model.Contact) *model.Contact); ok {
		r0 = rf(contact)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Contact) error); ok {
		r1 = rf(contact)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *ContactRepository) Delete(id int64) error {
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
func (_m *ContactRepository) Detail(id int64) (*model.Contact, error) {
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
func (_m *ContactRepository) List() ([]model.Contact, error) {
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

// Update provides a mock function with given fields: id, contact
func (_m *ContactRepository) Update(id int64, contact *model.Contact) (*model.Contact, error) {
	ret := _m.Called(id, contact)

	var r0 *model.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, *model.Contact) (*model.Contact, error)); ok {
		return rf(id, contact)
	}
	if rf, ok := ret.Get(0).(func(int64, *model.Contact) *model.Contact); ok {
		r0 = rf(id, contact)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, *model.Contact) error); ok {
		r1 = rf(id, contact)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewContactRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewContactRepository creates a new instance of ContactRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewContactRepository(t mockConstructorTestingTNewContactRepository) *ContactRepository {
	mock := &ContactRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
