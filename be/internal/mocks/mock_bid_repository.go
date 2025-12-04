package mocks

import (
	"milestone3/be/internal/entity"
	"milestone3/be/internal/repository"
	"time"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBidRedisRepository is a mock of BidRedisRepository interface.
type MockBidRedisRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBidRedisRepositoryMockRecorder
}

// MockBidRedisRepositoryMockRecorder is the mock recorder for MockBidRedisRepository.
type MockBidRedisRepositoryMockRecorder struct {
	mock *MockBidRedisRepository
}

// NewMockBidRedisRepository creates a new mock instance.
func NewMockBidRedisRepository(ctrl *gomock.Controller) *MockBidRedisRepository {
	mock := &MockBidRedisRepository{ctrl: ctrl}
	mock.recorder = &MockBidRedisRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBidRedisRepository) EXPECT() *MockBidRedisRepositoryMockRecorder {
	return m.recorder
}

// DeleteKey mocks base method.
func (m *MockBidRedisRepository) DeleteKey(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKey", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteKey indicates an expected call of DeleteKey.
func (mr *MockBidRedisRepositoryMockRecorder) DeleteKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKey", reflect.TypeOf((*MockBidRedisRepository)(nil).DeleteKey), key)
}

// GetBidByKey mocks base method.
func (m *MockBidRedisRepository) GetBidByKey(key string) (repository.BidEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBidByKey", key)
	ret0, _ := ret[0].(repository.BidEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBidByKey indicates an expected call of GetBidByKey.
func (mr *MockBidRedisRepositoryMockRecorder) GetBidByKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBidByKey", reflect.TypeOf((*MockBidRedisRepository)(nil).GetBidByKey), key)
}

// GetBidHistory mocks base method.
func (m *MockBidRedisRepository) GetBidHistory(sessionID, itemID int64, limit int64) ([]repository.BidEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBidHistory", sessionID, itemID, limit)
	ret0, _ := ret[0].([]repository.BidEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBidHistory indicates an expected call of GetBidHistory.
func (mr *MockBidRedisRepositoryMockRecorder) GetBidHistory(sessionID, itemID, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBidHistory", reflect.TypeOf((*MockBidRedisRepository)(nil).GetBidHistory), sessionID, itemID, limit)
}

// GetHighestBid mocks base method.
func (m *MockBidRedisRepository) GetHighestBid(sessionID, itemID int64) (float64, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHighestBid", sessionID, itemID)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetHighestBid indicates an expected call of GetHighestBid.
func (mr *MockBidRedisRepositoryMockRecorder) GetHighestBid(sessionID, itemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHighestBid", reflect.TypeOf((*MockBidRedisRepository)(nil).GetHighestBid), sessionID, itemID)
}

// GetSessionEndTime mocks base method.
func (m *MockBidRedisRepository) GetSessionEndTime(sessionID int64) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionEndTime", sessionID)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionEndTime indicates an expected call of GetSessionEndTime.
func (mr *MockBidRedisRepositoryMockRecorder) GetSessionEndTime(sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionEndTime", reflect.TypeOf((*MockBidRedisRepository)(nil).GetSessionEndTime), sessionID)
}

// ScanKeys mocks base method.
func (m *MockBidRedisRepository) ScanKeys(pattern string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanKeys", pattern)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScanKeys indicates an expected call of ScanKeys.
func (mr *MockBidRedisRepositoryMockRecorder) ScanKeys(pattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanKeys", reflect.TypeOf((*MockBidRedisRepository)(nil).ScanKeys), pattern)
}

// SetHighestBid mocks base method.
func (m *MockBidRedisRepository) SetHighestBid(sessionID, itemID int64, amount float64, userID int64, sessionEndTime time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHighestBid", sessionID, itemID, amount, userID, sessionEndTime)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHighestBid indicates an expected call of SetHighestBid.
func (mr *MockBidRedisRepositoryMockRecorder) SetHighestBid(sessionID, itemID, amount, userID, sessionEndTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHighestBid", reflect.TypeOf((*MockBidRedisRepository)(nil).SetHighestBid), sessionID, itemID, amount, userID, sessionEndTime)
}

// MockBidRepository is a mock of BidRepository interface.
type MockBidRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBidRepositoryMockRecorder
}

// MockBidRepositoryMockRecorder is the mock recorder for MockBidRepository.
type MockBidRepositoryMockRecorder struct {
	mock *MockBidRepository
}

// NewMockBidRepository creates a new mock instance.
func NewMockBidRepository(ctrl *gomock.Controller) *MockBidRepository {
	mock := &MockBidRepository{ctrl: ctrl}
	mock.recorder = &MockBidRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBidRepository) EXPECT() *MockBidRepositoryMockRecorder {
	return m.recorder
}

// SaveFinalBid mocks base method.
func (m *MockBidRepository) SaveFinalBid(bid *entity.Bid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFinalBid", bid)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveFinalBid indicates an expected call of SaveFinalBid.
func (mr *MockBidRepositoryMockRecorder) SaveFinalBid(bid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFinalBid", reflect.TypeOf((*MockBidRepository)(nil).SaveFinalBid), bid)
}