// Code generated by mockery v2.10.1. DO NOT EDIT.

package repository

import mock "github.com/stretchr/testify/mock"

// MockRollbackableRepository is an autogenerated mock type for the RollbackableRepository type
type MockRollbackableRepository struct {
	mock.Mock
}

// Rollback provides a mock function with given fields:
func (_m *MockRollbackableRepository) Rollback() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
