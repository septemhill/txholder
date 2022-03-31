// Code generated by mockery v2.10.1. DO NOT EDIT.

package repository

import mock "github.com/stretchr/testify/mock"

// MockTransactionHolderRepository is an autogenerated mock type for the TransactionHolderRepository type
type MockTransactionHolderRepository struct {
	mock.Mock
}

// Commit provides a mock function with given fields:
func (_m *MockTransactionHolderRepository) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rollback provides a mock function with given fields:
func (_m *MockTransactionHolderRepository) Rollback() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}