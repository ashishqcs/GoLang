// Code generated by mockery v2.30.16. DO NOT EDIT.

package mock

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GetAllMovies provides a mock function with given fields: ctx
func (_m *Service) GetAllMovies(ctx *gin.Context) {
	_m.Called(ctx)
}

// GetMovie provides a mock function with given fields: ctx, id
func (_m *Service) GetMovie(ctx *gin.Context, id string) {
	_m.Called(ctx, id)
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
