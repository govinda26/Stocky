package controllers

import (
	"stocky/initializers"
	"stocky/models"
	"time"

	"github.com/gin-gonic/gin"
)

func Stats(c *gin.Context) {
	//get user id from params
	id := c.Param("userId")

	//todays total
	today := time.Now().Truncate(24 * time.Hour)

	//initialise StockTotal
	type StockTotal struct {
		StockSymbol string
		TotalShares float64
	}

	//initialise todayTotals
	var todayTotals []StockTotal

	initializers.DB.Model(&models.RewardEvent{}).Select("stock_symbol, SUM(shares) as total_shares").Where("user_id = ? AND received_at >= ?", id, today).Group("stock_symbol").Scan(&todayTotals)

	//portfolio value
	var holdings []models.Holding
	initializers.DB.Where("user_id = ?", id).Find(&holdings)

	totalValue := 0.0
	for _, h := range holdings {
		price := generateRandomPrice(h.StockSymbol)
		totalValue = totalValue + price*h.TotalShares
	}

	//Round INR
	totalValue = float64(int(totalValue*100)) / 100

	//response
	c.JSON(200, gin.H{
		"todaysTotal":    todayTotals,
		"portfolioValue": totalValue,
	})
}
