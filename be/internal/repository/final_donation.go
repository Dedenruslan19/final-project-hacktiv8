package repository

import (
	"milestone3/be/internal/entity"

	"gorm.io/gorm"
)

type FinalDonationRepository interface {
	GetAllFinalDonations() ([]entity.FinalDonation, error)
	GetAllFinalDonationsByUserID(userID int) ([]entity.FinalDonation, error)
}

type finalDonationRepository struct {
	db *gorm.DB
}

func NewFinalDonationRepository(db *gorm.DB) FinalDonationRepository {
	return &finalDonationRepository{db: db}
}

// Return final_donations where the related donation has status = 'verified_for_donation'
func (r *finalDonationRepository) GetAllFinalDonations() ([]entity.FinalDonation, error) {
	var finalDonations []entity.FinalDonation
	err := r.db.
		Joins("JOIN donations d ON d.id = final_donations.donation_id").
		Where("d.status = ?", "verified_for_donation").
		Find(&finalDonations).Error
	return finalDonations, err
}

// Return final_donations for a user by joining donations and filtering by donation.user_id
func (r *finalDonationRepository) GetAllFinalDonationsByUserID(userID int) ([]entity.FinalDonation, error) {
	var finalDonations []entity.FinalDonation
	err := r.db.
		Joins("JOIN donations d ON d.id = final_donations.donation_id").
		Where("d.user_id = ? AND d.status = ?", userID, "verified_for_donation").
		Find(&finalDonations).Error
	return finalDonations, err
}
