package dto

type UpdateNotesDTO struct {
	DonationID uint   `json:"donation_id" validate:"required"`
	Notes      string `json:"notes" validate:"required"`
}
