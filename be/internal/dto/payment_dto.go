package dto

type PaymentRequest struct {
	UserId int `json:"user_id"`
	AuctionItemId float64 `json:"auction_item_id"`
	Amount int `json:"amount" validate:"required"`
}

type PaymentResponse struct {
	PaymentLinkUrl string `json:"payment_link_url"`
	TransactionId string `json:"transaction_id"`
	ExpiryTime string `json:"expiry_time"`
}

type CheckPaymentStatusResponse struct {
	
}