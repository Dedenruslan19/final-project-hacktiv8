package scheduler

import (
	"log/slog"
	"milestone3/be/internal/service"
	"time"

	"github.com/go-co-op/gocron"
)

type BidScheduler struct {
	bidSvc     service.BidService
	auctionSvc service.AuctionItemService
	logger     *slog.Logger
}

func NewBidScheduler(bidService service.BidService, auctionService service.AuctionItemService, logger *slog.Logger) *BidScheduler {
	return &BidScheduler{
		bidSvc:     bidService,
		auctionSvc: auctionService,
		logger:     logger,
	}
}

func (s *BidScheduler) Start() {
	// set as local time
	scheduler := gocron.NewScheduler(time.Local)

	// Check and start scheduled auctions every 1 minute
	_, err := scheduler.Every(1).Minute().Do(func() {
		s.logger.Info("Checking for scheduled auctions to start...")
		if err := s.auctionSvc.CheckAndStartScheduledItems(); err != nil {
			s.logger.Error("Failed to check/start scheduled items", "error", err)
		}
	})

	if err != nil {
		s.logger.Error("Failed to schedule auction auto-start", "error", err)
	}

	// save expired sessions to DB every 1 minute
	_, err = scheduler.Every(1).Minute().Do(func() {
		s.logger.Info("Running expired sessions cleanup...")
		if syncErr := s.bidSvc.SaveKeyToDB(); syncErr != nil {
			s.logger.Error("Failed to save expired sessions", "error", syncErr)
		}
		// Also close items without bids
		if closeErr := s.bidSvc.CloseExpiredItemsWithoutBids(); closeErr != nil {
			s.logger.Error("Failed to close items without bids", "error", closeErr)
		}
	})

	if err != nil {
		s.logger.Error("Failed to schedule bid sync", "error", err)
		return
	}

	// delete key value at 12 AM daily
	_, err = scheduler.Every(1).Day().At("00:00").Do(func() {
		s.logger.Info("Running midnight Redis cleanup...")
		if cleanupErr := s.bidSvc.DeleteKeyValue(); cleanupErr != nil {
			s.logger.Error("Failed to cleanup Redis at midnight", "error", cleanupErr)
		}
	})

	if err != nil {
		s.logger.Error("Failed to schedule midnight cleanup", "error", err)
		return
	}

	scheduler.StartAsync()
	s.logger.Info("Bid scheduler started")
	s.logger.Info("- Auto-start auctions: every 1 minute")
	s.logger.Info("- Sync to DB: every 1 minute")
	s.logger.Info("- Redis cleanup: daily at 00:00")
}
