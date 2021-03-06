// Code generated by MockGen. DO NOT EDIT.
// Source: .\moduleFetcher.go

// Package moduleFetcher is a generated GoMock package.
package moduleFetcher

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockModuleFetcher is a mock of ModuleFetcher interface.
type MockModuleFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockModuleFetcherMockRecorder
}

// MockModuleFetcherMockRecorder is the mock recorder for MockModuleFetcher.
type MockModuleFetcherMockRecorder struct {
	mock *MockModuleFetcher
}

// NewMockModuleFetcher creates a new mock instance.
func NewMockModuleFetcher(ctrl *gomock.Controller) *MockModuleFetcher {
	mock := &MockModuleFetcher{ctrl: ctrl}
	mock.recorder = &MockModuleFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModuleFetcher) EXPECT() *MockModuleFetcherMockRecorder {
	return m.recorder
}

// GetModule mocks base method.
func (m *MockModuleFetcher) GetModule(moduleName string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModule", moduleName)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModule indicates an expected call of GetModule.
func (mr *MockModuleFetcherMockRecorder) GetModule(moduleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModule", reflect.TypeOf((*MockModuleFetcher)(nil).GetModule), moduleName)
}
