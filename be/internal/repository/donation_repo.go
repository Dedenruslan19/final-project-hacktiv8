package repository

import (
	"milestone3/be/internal/entity"

	"gorm.io/gorm"
)

type DonationRepo interface {
	CreateDonation(donation entity.Donation) error
	GetDonationByID(id uint) (entity.Donation, error)
	UpdateDonation(donation entity.Donation) error
	DeleteDonation(id uint) error

	// Admin-only or filtered queries
	GetAllDonations() ([]entity.Donation, error)
	GetDonationsByUserID(userID uint) ([]entity.Donation, error)

	PatchDonation(donation entity.Donation) error
}

type donationRepo struct {
	db *gorm.DB
}

func NewDonationRepo(db *gorm.DB) DonationRepo {
	return &donationRepo{db: db}
}

func (r *donationRepo) CreateDonation(donation entity.Donation) error {
	return r.db.Create(&donation).Error
}

func (r *donationRepo) GetAllDonations() ([]entity.Donation, error) {
	var donations []entity.Donation
	err := r.db.Find(&donations).Error
	return donations, err
}

func (r *donationRepo) GetDonationsByUserID(userID uint) ([]entity.Donation, error) {
	var donations []entity.Donation
	err := r.db.Where("user_id = ?", userID).Find(&donations).Error
	return donations, err
}

func (r *donationRepo) GetDonationByID(id uint) (entity.Donation, error) {
	var donation entity.Donation
	err := r.db.First(&donation, id).Error
	return donation, err
}

func (r *donationRepo) UpdateDonation(donation entity.Donation) error {
	return r.db.Save(&donation).Error
}

func (r *donationRepo) DeleteDonation(id uint) error {
	return r.db.Delete(&entity.Donation{}, id).Error
}

func (r *donationRepo) PatchDonation(donation entity.Donation) error {
	return r.db.Model(&entity.Donation{}).Where("id = ?", donation.ID).Updates(donation).Error
}
