package dto

type PaymentRequest struct {
	OrderId int `json:"order_id"`
	GrossAmount float64 `json:"gross_amount" validate:"required,gte=1"`
	Name string `json:"name" validate:"required"`
	NoHp string `json:"no_hp" validate:"required,gte=5"`
	Email string `json:"email" validate:"required,email"`
}

type PaymentResponse struct {
	PaymentLinkUrl string `json:"payment_link_url"`
	TransactionId string `json:"transaction_id"`
	ExpiryTime string `json:"expiry_time"`
}