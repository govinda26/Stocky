package models

import "gorm.io/gorm"

type Holding struct {
	gorm.Model
	UserId      uint
	StockSymbol string
	TotalShares float64
}
