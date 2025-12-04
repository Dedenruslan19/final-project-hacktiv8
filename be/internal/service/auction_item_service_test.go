package service

import (
	"errors"
	"log/slog"
	"os"
	"testing"

	"milestone3/be/internal/dto"
	"milestone3/be/internal/entity"
	"milestone3/be/internal/mocks"


	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuctionItemService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAuctionItemRepository(ctrl)
	mockAI := mocks.NewMockAIRepository(ctrl)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	
	auctionService := NewAuctionItemService(mockRepo, mockAI, logger)

	tests := []struct {
		name    string
		req     dto.AuctionItemDTO
		setup   func()
		wantErr bool
	}{
		{
			name: "successful auction item creation",
			req: dto.AuctionItemDTO{
				Title:       "Test Item",
				Category:    "Electronics",
				Description: "Test description",
			},
			setup: func() {
				mockAI.EXPECT().EstimateStartingPrice(gomock.Any()).Return(float64(100), nil)
				mockRepo.EXPECT().Create(gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "AI estimation fails, fallback to default",
			req: dto.AuctionItemDTO{
				Title:       "Test Item",
				Category:    "Electronics",
				Description: "Test description",
			},
			setup: func() {
				mockAI.EXPECT().EstimateStartingPrice(gomock.Any()).Return(float64(0), errors.New("AI error"))
				mockRepo.EXPECT().Create(gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "repository create error",
			req: dto.AuctionItemDTO{
				Title:       "Test Item",
				Category:    "Electronics",
				Description: "Test description",
			},
			setup: func() {
				mockAI.EXPECT().EstimateStartingPrice(gomock.Any()).Return(float64(100), nil)
				mockRepo.EXPECT().Create(gomock.Any()).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			
			result, err := auctionService.Create(&tt.req)
			
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, "Test Item", result.Title)
			}
		})
	}
}

func TestAuctionItemService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAuctionItemRepository(ctrl)
	mockAI := mocks.NewMockAIRepository(ctrl)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	
	auctionService := NewAuctionItemService(mockRepo, mockAI, logger)

	tests := []struct {
		name    string
		setup   func()
		wantErr bool
	}{
		{
			name: "successful get all items",
			setup: func() {
				items := []entity.AuctionItem{
					{ID: 1, Title: "Item 1", StartingPrice: 100},
					{ID: 2, Title: "Item 2", StartingPrice: 200},
				}
				mockRepo.EXPECT().GetAll().Return(items, nil)
			},
			wantErr: false,
		},
		{
			name: "repository error",
			setup: func() {
				mockRepo.EXPECT().GetAll().Return(nil, errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			
			result, err := auctionService.GetAll()
			
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Len(t, result, 2)
			}
		})
	}
}

func TestAuctionItemService_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAuctionItemRepository(ctrl)
	mockAI := mocks.NewMockAIRepository(ctrl)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	
	auctionService := NewAuctionItemService(mockRepo, mockAI, logger)

	tests := []struct {
		name    string
		id      int64
		setup   func()
		wantErr bool
	}{
		{
			name: "successful get by id",
			id:   1,
			setup: func() {
				item := &entity.AuctionItem{ID: 1, Title: "Test Item", StartingPrice: 100}
				mockRepo.EXPECT().GetByID(int64(1)).Return(item, nil)
			},
			wantErr: false,
		},
		{
			name: "item not found",
			id:   999,
			setup: func() {
				mockRepo.EXPECT().GetByID(int64(999)).Return(nil, errors.New("not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			
			result, err := auctionService.GetByID(tt.id)
			
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, int64(1), result.ID)
			}
		})
	}
}