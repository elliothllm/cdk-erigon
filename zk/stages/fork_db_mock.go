// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ledgerwatch/erigon/zk/stages (interfaces: ForkDb)
//
// Generated by this command:
//
//	mockgen -typed=true -destination=./fork_db_mock.go -package=stages . ForkDb
//

// Package stages is a generated GoMock package.
package stages

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockForkDb is a mock of ForkDb interface.
type MockForkDb struct {
	ctrl     *gomock.Controller
	recorder *MockForkDbMockRecorder
}

// MockForkDbMockRecorder is the mock recorder for MockForkDb.
type MockForkDbMockRecorder struct {
	mock *MockForkDb
}

// NewMockForkDb creates a new mock instance.
func NewMockForkDb(ctrl *gomock.Controller) *MockForkDb {
	mock := &MockForkDb{ctrl: ctrl}
	mock.recorder = &MockForkDbMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockForkDb) EXPECT() *MockForkDbMockRecorder {
	return m.recorder
}

// GetAllForkHistory mocks base method.
func (m *MockForkDb) GetAllForkHistory() ([]uint64, []uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllForkHistory")
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].([]uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllForkHistory indicates an expected call of GetAllForkHistory.
func (mr *MockForkDbMockRecorder) GetAllForkHistory() *MockForkDbGetAllForkHistoryCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllForkHistory", reflect.TypeOf((*MockForkDb)(nil).GetAllForkHistory))
	return &MockForkDbGetAllForkHistoryCall{Call: call}
}

// MockForkDbGetAllForkHistoryCall wrap *gomock.Call
type MockForkDbGetAllForkHistoryCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockForkDbGetAllForkHistoryCall) Return(arg0, arg1 []uint64, arg2 error) *MockForkDbGetAllForkHistoryCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockForkDbGetAllForkHistoryCall) Do(f func() ([]uint64, []uint64, error)) *MockForkDbGetAllForkHistoryCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockForkDbGetAllForkHistoryCall) DoAndReturn(f func() ([]uint64, []uint64, error)) *MockForkDbGetAllForkHistoryCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetForkId mocks base method.
func (m *MockForkDb) GetForkId(arg0 uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForkId", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetForkId indicates an expected call of GetForkId.
func (mr *MockForkDbMockRecorder) GetForkId(arg0 any) *MockForkDbGetForkIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForkId", reflect.TypeOf((*MockForkDb)(nil).GetForkId), arg0)
	return &MockForkDbGetForkIdCall{Call: call}
}

// MockForkDbGetForkIdCall wrap *gomock.Call
type MockForkDbGetForkIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockForkDbGetForkIdCall) Return(arg0 uint64, arg1 error) *MockForkDbGetForkIdCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockForkDbGetForkIdCall) Do(f func(uint64) (uint64, error)) *MockForkDbGetForkIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockForkDbGetForkIdCall) DoAndReturn(f func(uint64) (uint64, error)) *MockForkDbGetForkIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetLatestForkHistory mocks base method.
func (m *MockForkDb) GetLatestForkHistory() (uint64, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestForkHistory")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetLatestForkHistory indicates an expected call of GetLatestForkHistory.
func (mr *MockForkDbMockRecorder) GetLatestForkHistory() *MockForkDbGetLatestForkHistoryCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestForkHistory", reflect.TypeOf((*MockForkDb)(nil).GetLatestForkHistory))
	return &MockForkDbGetLatestForkHistoryCall{Call: call}
}

// MockForkDbGetLatestForkHistoryCall wrap *gomock.Call
type MockForkDbGetLatestForkHistoryCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockForkDbGetLatestForkHistoryCall) Return(arg0, arg1 uint64, arg2 error) *MockForkDbGetLatestForkHistoryCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockForkDbGetLatestForkHistoryCall) Do(f func() (uint64, uint64, error)) *MockForkDbGetLatestForkHistoryCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockForkDbGetLatestForkHistoryCall) DoAndReturn(f func() (uint64, uint64, error)) *MockForkDbGetLatestForkHistoryCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WriteForkId mocks base method.
func (m *MockForkDb) WriteForkId(arg0, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteForkId", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteForkId indicates an expected call of WriteForkId.
func (mr *MockForkDbMockRecorder) WriteForkId(arg0, arg1 any) *MockForkDbWriteForkIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteForkId", reflect.TypeOf((*MockForkDb)(nil).WriteForkId), arg0, arg1)
	return &MockForkDbWriteForkIdCall{Call: call}
}

// MockForkDbWriteForkIdCall wrap *gomock.Call
type MockForkDbWriteForkIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockForkDbWriteForkIdCall) Return(arg0 error) *MockForkDbWriteForkIdCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockForkDbWriteForkIdCall) Do(f func(uint64, uint64) error) *MockForkDbWriteForkIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockForkDbWriteForkIdCall) DoAndReturn(f func(uint64, uint64) error) *MockForkDbWriteForkIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WriteForkIdBlockOnce mocks base method.
func (m *MockForkDb) WriteForkIdBlockOnce(arg0, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteForkIdBlockOnce", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteForkIdBlockOnce indicates an expected call of WriteForkIdBlockOnce.
func (mr *MockForkDbMockRecorder) WriteForkIdBlockOnce(arg0, arg1 any) *MockForkDbWriteForkIdBlockOnceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteForkIdBlockOnce", reflect.TypeOf((*MockForkDb)(nil).WriteForkIdBlockOnce), arg0, arg1)
	return &MockForkDbWriteForkIdBlockOnceCall{Call: call}
}

// MockForkDbWriteForkIdBlockOnceCall wrap *gomock.Call
type MockForkDbWriteForkIdBlockOnceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockForkDbWriteForkIdBlockOnceCall) Return(arg0 error) *MockForkDbWriteForkIdBlockOnceCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockForkDbWriteForkIdBlockOnceCall) Do(f func(uint64, uint64) error) *MockForkDbWriteForkIdBlockOnceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockForkDbWriteForkIdBlockOnceCall) DoAndReturn(f func(uint64, uint64) error) *MockForkDbWriteForkIdBlockOnceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
