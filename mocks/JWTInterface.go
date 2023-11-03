// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	jwt "github.com/golang-jwt/jwt/v5"
	mock "github.com/stretchr/testify/mock"
)

// JWTInterface is an autogenerated mock type for the JWTInterface type
type JWTInterface struct {
	mock.Mock
}

// ExtractToken provides a mock function with given fields: token
func (_m *JWTInterface) ExtractToken(token *jwt.Token) uint {
	ret := _m.Called(token)

	var r0 uint
	if rf, ok := ret.Get(0).(func(*jwt.Token) uint); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// GenerateJWT provides a mock function with given fields: userID
func (_m *JWTInterface) GenerateJWT(userID uint) string {
	ret := _m.Called(userID)

	var r0 string
	if rf, ok := ret.Get(0).(func(uint) string); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GenerateToken provides a mock function with given fields: id
func (_m *JWTInterface) GenerateToken(id uint) string {
	ret := _m.Called(id)

	var r0 string
	if rf, ok := ret.Get(0).(func(uint) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewJWTInterface creates a new instance of JWTInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJWTInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *JWTInterface {
	mock := &JWTInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
