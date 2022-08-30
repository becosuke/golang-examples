// Code generated by MockGen. DO NOT EDIT.
// Source: internal/drivers/syncmap/syncmap.go

// Package syncmap is a generated GoMock package.
package syncmap

import (
	reflect "reflect"

	syncmap "github.com/becosuke/golang-examples/kvstore/internal/drivers/syncmap"
	gomock "github.com/golang/mock/gomock"
)

// MockSyncmap is a mock of Syncmap interface.
type MockSyncmap struct {
	ctrl     *gomock.Controller
	recorder *MockSyncmapMockRecorder
}

// MockSyncmapMockRecorder is the mock recorder for MockSyncmap.
type MockSyncmapMockRecorder struct {
	mock *MockSyncmap
}

// NewMockSyncmap creates a new mock instance.
func NewMockSyncmap(ctrl *gomock.Controller) *MockSyncmap {
	mock := &MockSyncmap{ctrl: ctrl}
	mock.recorder = &MockSyncmapMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyncmap) EXPECT() *MockSyncmapMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockSyncmap) Delete(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSyncmapMockRecorder) Delete(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSyncmap)(nil).Delete), key)
}

// Load mocks base method.
func (m *MockSyncmap) Load(key string) (*syncmap.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", key)
	ret0, _ := ret[0].(*syncmap.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load.
func (mr *MockSyncmapMockRecorder) Load(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockSyncmap)(nil).Load), key)
}

// LoadOrStore mocks base method.
func (m *MockSyncmap) LoadOrStore(message *syncmap.Message) (*syncmap.Message, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadOrStore", message)
	ret0, _ := ret[0].(*syncmap.Message)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoadOrStore indicates an expected call of LoadOrStore.
func (mr *MockSyncmapMockRecorder) LoadOrStore(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadOrStore", reflect.TypeOf((*MockSyncmap)(nil).LoadOrStore), message)
}

// Store mocks base method.
func (m *MockSyncmap) Store(message *syncmap.Message) (*syncmap.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", message)
	ret0, _ := ret[0].(*syncmap.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store.
func (mr *MockSyncmapMockRecorder) Store(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockSyncmap)(nil).Store), message)
}
