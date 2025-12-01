package entity

import "time"

type FinalDonation struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	DonationID uint      `gorm:"not null" json:"donation_id"`
	Donation   Donation  `gorm:"foreignKey:DonationID" json:"donation,omitempty"` // preload-able
	Notes      string    `gorm:"type:text" json:"notes"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}
