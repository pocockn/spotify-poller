// Code generated by MockGen. DO NOT EDIT.
// Source: storer.go

// Package mock_internals is a generated GoMock package.
package mock_internals

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/pocockn/recs-api/models"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockStore) Create(rec *models.Rec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", rec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockStoreMockRecorder) Create(rec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStore)(nil).Create), rec)
}

// FetchAll mocks base method
func (m *MockStore) FetchAll() (models.Recs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAll")
	ret0, _ := ret[0].(models.Recs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAll indicates an expected call of FetchAll
func (mr *MockStoreMockRecorder) FetchAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAll", reflect.TypeOf((*MockStore)(nil).FetchAll))
}
