//go:generate mockgen -source=../service/admin_service.go -destination=mock_admin_repository.go
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAdminRepository is a mock of AdminRepository interface.
type MockAdminRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAdminRepositoryMockRecorder
}

// MockAdminRepositoryMockRecorder is the mock recorder for MockAdminRepository.
type MockAdminRepositoryMockRecorder struct {
	mock *MockAdminRepository
}

// NewMockAdminRepository creates a new mock instance.
func NewMockAdminRepository(ctrl *gomock.Controller) *MockAdminRepository {
	mock := &MockAdminRepository{ctrl: ctrl}
	mock.recorder = &MockAdminRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminRepository) EXPECT() *MockAdminRepositoryMockRecorder {
	return m.recorder
}

// CountArticle mocks base method.
func (m *MockAdminRepository) CountArticle() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountArticle")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountArticle indicates an expected call of CountArticle.
func (mr *MockAdminRepositoryMockRecorder) CountArticle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountArticle", reflect.TypeOf((*MockAdminRepository)(nil).CountArticle))
}

// CountAuction mocks base method.
func (m *MockAdminRepository) CountAuction() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountAuction")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountAuction indicates an expected call of CountAuction.
func (mr *MockAdminRepositoryMockRecorder) CountAuction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountAuction", reflect.TypeOf((*MockAdminRepository)(nil).CountAuction))
}

// CountDonation mocks base method.
func (m *MockAdminRepository) CountDonation() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountDonation")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountDonation indicates an expected call of CountDonation.
func (mr *MockAdminRepositoryMockRecorder) CountDonation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountDonation", reflect.TypeOf((*MockAdminRepository)(nil).CountDonation))
}

// CountPayment mocks base method.
func (m *MockAdminRepository) CountPayment() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountPayment")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountPayment indicates an expected call of CountPayment.
func (mr *MockAdminRepositoryMockRecorder) CountPayment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountPayment", reflect.TypeOf((*MockAdminRepository)(nil).CountPayment))
}