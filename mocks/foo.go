// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nonki/golang/foo (interfaces: Issuer,Container)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	foo "github.com/nonki/golang/foo"
	reflect "reflect"
)

// MockIssuer is a mock of Issuer interface
type MockIssuer struct {
	ctrl     *gomock.Controller
	recorder *MockIssuerMockRecorder
}

// MockIssuerMockRecorder is the mock recorder for MockIssuer
type MockIssuerMockRecorder struct {
	mock *MockIssuer
}

// NewMockIssuer creates a new mock instance
func NewMockIssuer(ctrl *gomock.Controller) *MockIssuer {
	mock := &MockIssuer{ctrl: ctrl}
	mock.recorder = &MockIssuerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIssuer) EXPECT() *MockIssuerMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockIssuer) Get() foo.Container {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(foo.Container)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockIssuerMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIssuer)(nil).Get))
}

// MockContainer is a mock of Container interface
type MockContainer struct {
	ctrl     *gomock.Controller
	recorder *MockContainerMockRecorder
}

// MockContainerMockRecorder is the mock recorder for MockContainer
type MockContainerMockRecorder struct {
	mock *MockContainer
}

// NewMockContainer creates a new mock instance
func NewMockContainer(ctrl *gomock.Controller) *MockContainer {
	mock := &MockContainer{ctrl: ctrl}
	mock.recorder = &MockContainerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContainer) EXPECT() *MockContainerMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockContainer) Get() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(string)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockContainerMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockContainer)(nil).Get))
}