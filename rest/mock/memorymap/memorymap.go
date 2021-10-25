// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructure/memorymap/memorymap.go

// Package memorymap is a generated GoMock package.
package memorymap

import (
	reflect "reflect"

	memorymap "github.com/becosuke/golang-examples/rest/infrastructure/memorymap"
	gomock "github.com/golang/mock/gomock"
)

// MockMemoryMap is a mock of MemoryMap interface.
type MockMemoryMap struct {
	ctrl     *gomock.Controller
	recorder *MockMemoryMapMockRecorder
}

// MockMemoryMapMockRecorder is the mock recorder for MockMemoryMap.
type MockMemoryMapMockRecorder struct {
	mock *MockMemoryMap
}

// NewMockMemoryMap creates a new mock instance.
func NewMockMemoryMap(ctrl *gomock.Controller) *MockMemoryMap {
	mock := &MockMemoryMap{ctrl: ctrl}
	mock.recorder = &MockMemoryMapMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMemoryMap) EXPECT() *MockMemoryMapMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockMemoryMap) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMemoryMapMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMemoryMap)(nil).Delete), arg0)
}

// Load mocks base method.
func (m *MockMemoryMap) Load(arg0 string) (*memorymap.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", arg0)
	ret0, _ := ret[0].(*memorymap.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load.
func (mr *MockMemoryMapMockRecorder) Load(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockMemoryMap)(nil).Load), arg0)
}

// LoadOrStore mocks base method.
func (m *MockMemoryMap) LoadOrStore(arg0, arg1 string) (*memorymap.Message, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadOrStore", arg0, arg1)
	ret0, _ := ret[0].(*memorymap.Message)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoadOrStore indicates an expected call of LoadOrStore.
func (mr *MockMemoryMapMockRecorder) LoadOrStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadOrStore", reflect.TypeOf((*MockMemoryMap)(nil).LoadOrStore), arg0, arg1)
}

// Store mocks base method.
func (m *MockMemoryMap) Store(arg0, arg1 string) (*memorymap.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(*memorymap.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store.
func (mr *MockMemoryMapMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockMemoryMap)(nil).Store), arg0, arg1)
}
