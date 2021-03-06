// Code generated by MockGen. DO NOT EDIT.
// Source: authorize.go

// Package main is a generated GoMock package.
package main

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// authorize mocks base method.
func (m *MockAuthService) authorize(req *http.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "authorize", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// authorize indicates an expected call of authorize.
func (mr *MockAuthServiceMockRecorder) authorize(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "authorize", reflect.TypeOf((*MockAuthService)(nil).authorize), req)
}
