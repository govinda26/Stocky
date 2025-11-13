package models

import (
	"time"

	"gorm.io/gorm"
)

// RewardEvent is the table name
type RewardEvent struct {
	gorm.Model
	UserID      uint //userId
	StockSymbol string
	Shares      float64
	ReceivedAt  time.Time
	Source      string
	ExternalRef string `gorm:"unique"` //makes sure each reward has unique number so that user does not get same rewards again and again
}
