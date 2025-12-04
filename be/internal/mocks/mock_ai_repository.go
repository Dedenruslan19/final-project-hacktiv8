package mocks

import (
	"milestone3/be/internal/repository"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAIRepository is a mock of AIRepository interface.
type MockAIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAIRepositoryMockRecorder
}

// MockAIRepositoryMockRecorder is the mock recorder for MockAIRepository.
type MockAIRepositoryMockRecorder struct {
	mock *MockAIRepository
}

// NewMockAIRepository creates a new mock instance.
func NewMockAIRepository(ctrl *gomock.Controller) *MockAIRepository {
	mock := &MockAIRepository{ctrl: ctrl}
	mock.recorder = &MockAIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAIRepository) EXPECT() *MockAIRepositoryMockRecorder {
	return m.recorder
}

// EstimateStartingPrice mocks base method.
func (m *MockAIRepository) EstimateStartingPrice(req repository.PriceEstimationRequest) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateStartingPrice", req)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateStartingPrice indicates an expected call of EstimateStartingPrice.
func (mr *MockAIRepositoryMockRecorder) EstimateStartingPrice(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateStartingPrice", reflect.TypeOf((*MockAIRepository)(nil).EstimateStartingPrice), req)
}