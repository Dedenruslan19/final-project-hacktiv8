package service

import (
	"errors"
	"testing"

	"milestone3/be/internal/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAdminService_AdminDashboard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAdminRepository(ctrl)
	adminService := NewAdminService(mockRepo)

	tests := []struct {
		name    string
		setup   func()
		wantErr bool
	}{
		{
			name: "successful dashboard data retrieval",
			setup: func() {
				mockRepo.EXPECT().CountArticle().Return(int64(5), nil)
				mockRepo.EXPECT().CountDonation().Return(int64(10), nil)
				mockRepo.EXPECT().CountPayment().Return(int64(3), nil)
				mockRepo.EXPECT().CountAuction().Return(int64(7), nil)
			},
			wantErr: false,
		},
		{
			name: "article count error",
			setup: func() {
				mockRepo.EXPECT().CountArticle().Return(int64(0), errors.New("db error"))
			},
			wantErr: true,
		},
		{
			name: "donation count error",
			setup: func() {
				mockRepo.EXPECT().CountArticle().Return(int64(5), nil)
				mockRepo.EXPECT().CountDonation().Return(int64(0), errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			
			result, err := adminService.AdminDashboard()
			
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, int64(5), result.TotalArticle)
				assert.Equal(t, int64(10), result.TotalDonation)
				assert.Equal(t, int64(3), result.TotalPayment)
				assert.Equal(t, int64(7), result.TotalAuction)
			}
		})
	}
}