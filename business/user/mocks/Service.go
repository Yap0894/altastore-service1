// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	user "AltaStore/business/user"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: id, deleter
func (_m *Service) DeleteUser(id string, deleter string) error {
	ret := _m.Called(id, deleter)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(id, deleter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindUserByEmail provides a mock function with given fields: email
func (_m *Service) FindUserByEmail(email string) (*user.User, error) {
	ret := _m.Called(email)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(string) *user.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByEmailAndPassword provides a mock function with given fields: email, password
func (_m *Service) FindUserByEmailAndPassword(email string, password string) (*user.User, error) {
	ret := _m.Called(email, password)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(string, string) *user.User); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByID provides a mock function with given fields: id
func (_m *Service) FindUserByID(id string) (*user.User, error) {
	ret := _m.Called(id)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(string) *user.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertUser provides a mock function with given fields: insertUserSpec
func (_m *Service) InsertUser(insertUserSpec user.InsertUserSpec) error {
	ret := _m.Called(insertUserSpec)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.InsertUserSpec) error); ok {
		r0 = rf(insertUserSpec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: id, updateUserSpec, modifier
func (_m *Service) UpdateUser(id string, updateUserSpec user.UpdateUserSpec, modifier string) error {
	ret := _m.Called(id, updateUserSpec, modifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, user.UpdateUserSpec, string) error); ok {
		r0 = rf(id, updateUserSpec, modifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUserPassword provides a mock function with given fields: id, password, oldPassword, modifier
func (_m *Service) UpdateUserPassword(id string, password string, oldPassword string, modifier string) error {
	ret := _m.Called(id, password, oldPassword, modifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(id, password, oldPassword, modifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}