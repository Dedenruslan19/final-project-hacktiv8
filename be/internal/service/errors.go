package service

import "errors"

var (
	// User Errors
	ErrUserNotFound         = errors.New("user not found")
	ErrInvalidUser          = errors.New("invalid user data")
	ErrUserNotFoundID       = errors.New("user ID not found")
	ErrUserNotFoundName     = errors.New("user name not found")
	ErrUserNotFoundEmail    = errors.New("user email not found")
	ErrUserNotFoundPassword = errors.New("user password not found")
	// Payment Errors
	ErrPaymentNotFound       = errors.New("payment not found")
	ErrInvalidPayment        = errors.New("invalid payment data")
	ErrPaymentNotFoundID     = errors.New("payment ID not found")
	ErrPaymentNotFoundAmount = errors.New("payment amount not found")
	ErrPaymentNotFoundMethod = errors.New("payment method not found")
	ErrPaymentNotFoundStatus = errors.New("payment status not found")
	// Auction Errors
	ErrAuctionNotFound   = errors.New("auction not found")
	ErrInvalidAuction    = errors.New("invalid auction data")
	ErrAuctionNotFoundID = errors.New("auction ID not found")
	ErrSessionNotFoundID = errors.New("auction session ID not found")
	ErrInvalidDate       = errors.New("end time should be after start time")
	ErrInvalidTime       = errors.New("time must be in the future")
	ErrActiveSession     = errors.New("cannot modify an active auction session")
	ErrAuctionFinished   = errors.New("cannot update auction item with status 'finished'")
	ErrExpiredSession    = errors.New("cannot modify expired auction session")
	// Donation Errors
	ErrDonationNotFound          = errors.New("donation not found")
	ErrInvalidDonation           = errors.New("invalid donation data")
	ErrDonationNotFoundID        = errors.New("donation ID not found")
	ErrDonationNotFoundAmount    = errors.New("donation amount not found")
	ErrDonationNotFoundDonorName = errors.New("donor name not found")
	// Article Errors
	ErrArticleNotFound = errors.New("article not found")
	ErrInvalidArticle  = errors.New("invalid article data")
	// Bidding Errors
	ErrInvalidBidding       = errors.New("invalid bid amount")
	ErrBidTooLow            = errors.New("bid too low")
	ErrDuplicateBid         = errors.New("duplicate bid")
	ErrAlreadyHighestBidder = errors.New("you are already the highest bidder")
	// Final Donation Errors
	ErrFinalDonationNotFound   = errors.New("final donation not found")
	ErrFinalDonationNotFoundID = errors.New("final donation ID not found")
	// image Errors
	ErrImageNotFound   = errors.New("image not found")
	ErrSignedURLFailed = errors.New("signed URL generation failed")

	// Authorization / Generic Errors
	ErrUnauthorized = errors.New("unauthorized access")
	ErrForbidden    = errors.New("forbidden access")
)
