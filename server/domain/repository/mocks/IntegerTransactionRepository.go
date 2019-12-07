// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/CodFrm/cxmooc-tools/server/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// IntegerTransactionRepository is an autogenerated mock type for the IntegerTransactionRepository type
type IntegerTransactionRepository struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *IntegerTransactionRepository) Close() {
	_m.Called()
}

// Commit provides a mock function with given fields:
func (_m *IntegerTransactionRepository) Commit() {
	_m.Called()
}

// LockIntegral provides a mock function with given fields: token
func (_m *IntegerTransactionRepository) LockIntegral(token string) (*entity.IntegralEntity, error) {
	ret := _m.Called(token)

	var r0 *entity.IntegralEntity
	if rf, ok := ret.Get(0).(func(string) *entity.IntegralEntity); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.IntegralEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Rollback provides a mock function with given fields:
func (_m *IntegerTransactionRepository) Rollback() {
	_m.Called()
}

// Update provides a mock function with given fields: integral
func (_m *IntegerTransactionRepository) Update(integral *entity.IntegralEntity) error {
	ret := _m.Called(integral)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.IntegralEntity) error); ok {
		r0 = rf(integral)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
