package service

import (
	"errors"
	"log/slog"
	"os"
	"testing"
	"time"

	"milestone3/be/internal/entity"
	"milestone3/be/internal/mocks"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestBidService_PlaceBid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRedisRepo := mocks.NewMockBidRedisRepository(ctrl)
	mockBidRepo := mocks.NewMockBidRepository(ctrl)
	mockItemRepo := mocks.NewMockAuctionItemRepository(ctrl)
	mockSessionRepo := mocks.NewMockAuctionSessionRepository(ctrl)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	bidService := NewBidService(mockRedisRepo, mockBidRepo, mockItemRepo, mockSessionRepo, logger)

	tests := []struct {
		name           string
		sessionID      int64
		itemID         int64
		userID         int64
		amount         float64
		sessionEndTime time.Time
		setup          func()
		wantErr        bool
	}{
		{
			name:           "successful bid placement",
			sessionID:      1,
			itemID:         1,
			userID:         1,
			amount:         150.0,
			sessionEndTime: time.Now().Add(time.Hour),
			setup: func() {
				item := &entity.AuctionItem{
					ID:     1,
					Status: "ongoing",
				}
				mockItemRepo.EXPECT().GetByID(int64(1)).Return(item, nil)
				mockRedisRepo.EXPECT().GetHighestBid(int64(1), int64(1)).Return(100.0, int64(2), nil)
				mockRedisRepo.EXPECT().SetHighestBid(int64(1), int64(1), 150.0, int64(1), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name:           "invalid bid amount - zero",
			sessionID:      1,
			itemID:         1,
			userID:         1,
			amount:         0,
			sessionEndTime: time.Now().Add(time.Hour),
			setup:          func() {},
			wantErr:        true,
		},
		{
			name:           "item not found",
			sessionID:      1,
			itemID:         999,
			userID:         1,
			amount:         150.0,
			sessionEndTime: time.Now().Add(time.Hour),
			setup: func() {
				mockItemRepo.EXPECT().GetByID(int64(999)).Return(nil, errors.New("not found"))
			},
			wantErr: true,
		},
		{
			name:           "item not ongoing",
			sessionID:      1,
			itemID:         1,
			userID:         1,
			amount:         150.0,
			sessionEndTime: time.Now().Add(time.Hour),
			setup: func() {
				item := &entity.AuctionItem{
					ID:     1,
					Status: "finished",
				}
				mockItemRepo.EXPECT().GetByID(int64(1)).Return(item, nil)
			},
			wantErr: true,
		},
		{
			name:           "bid too low",
			sessionID:      1,
			itemID:         1,
			userID:         1,
			amount:         50.0,
			sessionEndTime: time.Now().Add(time.Hour),
			setup: func() {
				item := &entity.AuctionItem{
					ID:     1,
					Status: "ongoing",
				}
				mockItemRepo.EXPECT().GetByID(int64(1)).Return(item, nil)
				mockRedisRepo.EXPECT().GetHighestBid(int64(1), int64(1)).Return(100.0, int64(2), nil)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			err := bidService.PlaceBid(tt.sessionID, tt.itemID, tt.userID, tt.amount, tt.sessionEndTime)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestBidService_GetHighestBid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRedisRepo := mocks.NewMockBidRedisRepository(ctrl)
	mockBidRepo := mocks.NewMockBidRepository(ctrl)
	mockItemRepo := mocks.NewMockAuctionItemRepository(ctrl)
	mockSessionRepo := mocks.NewMockAuctionSessionRepository(ctrl)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	bidService := NewBidService(mockRedisRepo, mockBidRepo, mockItemRepo, mockSessionRepo, logger)

	tests := []struct {
		name      string
		sessionID int64
		itemID    int64
		setup     func()
		wantErr   bool
	}{
		{
			name:      "successful get highest bid",
			sessionID: 1,
			itemID:    1,
			setup: func() {
				item := &entity.AuctionItem{ID: 1}
				mockItemRepo.EXPECT().GetByID(int64(1)).Return(item, nil)
				mockRedisRepo.EXPECT().GetHighestBid(int64(1), int64(1)).Return(150.0, int64(1), nil)
			},
			wantErr: false,
		},
		{
			name:      "item not found",
			sessionID: 1,
			itemID:    999,
			setup: func() {
				mockItemRepo.EXPECT().GetByID(int64(999)).Return(nil, errors.New("not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			amount, bidder, err := bidService.GetHighestBid(tt.sessionID, tt.itemID)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, float64(0), amount)
				assert.Equal(t, int64(0), bidder)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, 150.0, amount)
				assert.Equal(t, int64(1), bidder)
			}
		})
	}
}
