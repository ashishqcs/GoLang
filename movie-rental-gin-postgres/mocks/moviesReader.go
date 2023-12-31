// Code generated by MockGen. DO NOT EDIT.
// Source: movieRentals/service (interfaces: MoviesReader)

// Package mocks is a generated GoMock package.
package mocks

import (
	model "movieRentals/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMoviesReader is a mock of MoviesReader interface.
type MockMoviesReader struct {
	ctrl     *gomock.Controller
	recorder *MockMoviesReaderMockRecorder
}

// MockMoviesReaderMockRecorder is the mock recorder for MockMoviesReader.
type MockMoviesReaderMockRecorder struct {
	mock *MockMoviesReader
}

// NewMockMoviesReader creates a new mock instance.
func NewMockMoviesReader(ctrl *gomock.Controller) *MockMoviesReader {
	mock := &MockMoviesReader{ctrl: ctrl}
	mock.recorder = &MockMoviesReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMoviesReader) EXPECT() *MockMoviesReaderMockRecorder {
	return m.recorder
}

// GetMovies mocks base method.
func (m *MockMoviesReader) GetMovies() (*model.Movies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovies")
	ret0, _ := ret[0].(*model.Movies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovies indicates an expected call of GetMovies.
func (mr *MockMoviesReaderMockRecorder) GetMovies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovies", reflect.TypeOf((*MockMoviesReader)(nil).GetMovies))
}
