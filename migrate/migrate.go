package main

import (
	"stocky/initializers"
	"stocky/models"
)

func init() {
	initializers.LoadEnvVariables() //gets port and db string from env
	initializers.ConnectToDB()      //gets DB connection variable
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.RewardEvent{}, &models.Holding{}, &models.StockPrice{}, &models.Ledger{})
}
