// Code generated by mockery v2.10.1. DO NOT EDIT.

package repository

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// MockConfig is an autogenerated mock type for the Config type
type MockConfig struct {
	mock.Mock
}

// Db provides a mock function with given fields: _a0
func (_m *MockConfig) Db(_a0 *gorm.DB) {
	_m.Called(_a0)
}

// Tx provides a mock function with given fields: _a0
func (_m *MockConfig) Tx(_a0 *gorm.DB) {
	_m.Called(_a0)
}
