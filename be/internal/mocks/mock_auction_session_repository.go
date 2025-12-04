package mocks

import (
	"milestone3/be/internal/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAuctionSessionRepository is a mock of AuctionSessionRepository interface.
type MockAuctionSessionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuctionSessionRepositoryMockRecorder
}

// MockAuctionSessionRepositoryMockRecorder is the mock recorder for MockAuctionSessionRepository.
type MockAuctionSessionRepositoryMockRecorder struct {
	mock *MockAuctionSessionRepository
}

// NewMockAuctionSessionRepository creates a new mock instance.
func NewMockAuctionSessionRepository(ctrl *gomock.Controller) *MockAuctionSessionRepository {
	mock := &MockAuctionSessionRepository{ctrl: ctrl}
	mock.recorder = &MockAuctionSessionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuctionSessionRepository) EXPECT() *MockAuctionSessionRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAuctionSessionRepository) Create(session *entity.AuctionSession) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", session)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAuctionSessionRepositoryMockRecorder) Create(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuctionSessionRepository)(nil).Create), session)
}

// Delete mocks base method.
func (m *MockAuctionSessionRepository) Delete(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAuctionSessionRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuctionSessionRepository)(nil).Delete), id)
}

// GetActiveSessions mocks base method.
func (m *MockAuctionSessionRepository) GetActiveSessions() ([]*entity.AuctionSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveSessions")
	ret0, _ := ret[0].([]*entity.AuctionSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveSessions indicates an expected call of GetActiveSessions.
func (mr *MockAuctionSessionRepositoryMockRecorder) GetActiveSessions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveSessions", reflect.TypeOf((*MockAuctionSessionRepository)(nil).GetActiveSessions))
}

// GetAll mocks base method.
func (m *MockAuctionSessionRepository) GetAll() ([]*entity.AuctionSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*entity.AuctionSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockAuctionSessionRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockAuctionSessionRepository)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockAuctionSessionRepository) GetByID(id int64) (*entity.AuctionSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*entity.AuctionSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockAuctionSessionRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockAuctionSessionRepository)(nil).GetByID), id)
}

// Update mocks base method.
func (m *MockAuctionSessionRepository) Update(session *entity.AuctionSession) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", session)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAuctionSessionRepositoryMockRecorder) Update(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAuctionSessionRepository)(nil).Update), session)
}