package models

import "gorm.io/gorm"

type Holding struct {
	gorm.Model
	UserID      uint
	StockSymbol string
	TotalShares float64
}
