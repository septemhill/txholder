// Code generated by mockery v2.10.1. DO NOT EDIT.

package repository

import mock "github.com/stretchr/testify/mock"

// MockTransactionalRepositoryCreator is an autogenerated mock type for the TransactionalRepositoryCreator type
type MockTransactionalRepositoryCreator struct {
	mock.Mock
}

// NewApplicationTxHolderRepository provides a mock function with given fields:
func (_m *MockTransactionalRepositoryCreator) NewApplicationTxHolderRepository() ApplicationTxHolderRepository {
	ret := _m.Called()

	var r0 ApplicationTxHolderRepository
	if rf, ok := ret.Get(0).(func() ApplicationTxHolderRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ApplicationTxHolderRepository)
		}
	}

	return r0
}

// NewUserTxHolderRepository provides a mock function with given fields:
func (_m *MockTransactionalRepositoryCreator) NewUserTxHolderRepository() UserTxHolderRepository {
	ret := _m.Called()

	var r0 UserTxHolderRepository
	if rf, ok := ret.Get(0).(func() UserTxHolderRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(UserTxHolderRepository)
		}
	}

	return r0
}
