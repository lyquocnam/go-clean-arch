// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import gin "github.com/gin-gonic/gin"

import mock "github.com/stretchr/testify/mock"

// CustomerHandler is an autogenerated mock type for the CustomerHandler type
type CustomerHandler struct {
	mock.Mock
}

// GetCustomerByIdHandler provides a mock function with given fields: c
func (_m *CustomerHandler) GetCustomerByIdHandler(c *gin.Context) {
	_m.Called(c)
}
