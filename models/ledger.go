package models

import "gorm.io/gorm"

type Ledger struct {
	gorm.Model
	EntryGroupId string //links entries from same transaction
	Account      string //"stock:name", cash=total share cost, fees=total brokerage to NSE/BSE.
	Debit        string
	Credit       string
	Shares       float64
	StockSymbol  string
}
