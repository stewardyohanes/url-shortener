package models

import "time"

type (
	ShortenRequest struct {
		URL string `json:"url" binding:"required,url"`
	}
)

type URL struct {
	ID          uint   `gorm:"primaryKey"`
	OriginalURL string `gorm:"not null"`
	ShortCode   string `gorm:"uniqueIndex;size:10"`
	VisitCount  int    `gorm:"default:0"`
	CreatedAt   time.Time
}
