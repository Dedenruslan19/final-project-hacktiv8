package mocks

import (
	"milestone3/be/internal/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSessionRedisRepository is a mock of SessionRedisRepository interface.
type MockSessionRedisRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSessionRedisRepositoryMockRecorder
}

// MockSessionRedisRepositoryMockRecorder is the mock recorder for MockSessionRedisRepository.
type MockSessionRedisRepositoryMockRecorder struct {
	mock *MockSessionRedisRepository
}

// NewMockSessionRedisRepository creates a new mock instance.
func NewMockSessionRedisRepository(ctrl *gomock.Controller) *MockSessionRedisRepository {
	mock := &MockSessionRedisRepository{ctrl: ctrl}
	mock.recorder = &MockSessionRedisRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionRedisRepository) EXPECT() *MockSessionRedisRepositoryMockRecorder {
	return m.recorder
}

// DeleteSession mocks base method.
func (m *MockSessionRedisRepository) DeleteSession(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockSessionRedisRepositoryMockRecorder) DeleteSession(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockSessionRedisRepository)(nil).DeleteSession), id)
}

// SetActiveSession mocks base method.
func (m *MockSessionRedisRepository) SetActiveSession(session entity.AuctionSession) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetActiveSession", session)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetActiveSession indicates an expected call of SetActiveSession.
func (mr *MockSessionRedisRepositoryMockRecorder) SetActiveSession(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetActiveSession", reflect.TypeOf((*MockSessionRedisRepository)(nil).SetActiveSession), session)
}