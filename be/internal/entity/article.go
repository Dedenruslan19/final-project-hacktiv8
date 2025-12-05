package entity

import (
	"time"
)

type Article struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"size:255;not null" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	Week      int       `json:"week"` // Format: YYYYMMDD (e.g., 20241204 for 04 Dec 2024)
	Image     string    `gorm:"type:text" json:"image"` // URL IMAGE (public bucket)
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
