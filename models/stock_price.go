package models

import (
	"time"

	"gorm.io/gorm"
)

type StockPrice struct {
	gorm.Model
	StockSymbol string
	PriceINR    float64
	PriceTime   time.Time
}
