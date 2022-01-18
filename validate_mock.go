// Code generated by MockGen. DO NOT EDIT.
// Source: validate.go

// Package main is a generated GoMock package.
package main

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockValidator is a mock of Validator interface.
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator.
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance.
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// validate mocks base method.
func (m *MockValidator) validate(appRequest *AppRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "validate", appRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// validate indicates an expected call of validate.
func (mr *MockValidatorMockRecorder) validate(appRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "validate", reflect.TypeOf((*MockValidator)(nil).validate), appRequest)
}