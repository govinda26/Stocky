package controllers

import (
	"stocky/initializers"
	"stocky/models"
	"time"

	"github.com/gin-gonic/gin"
)

func HistoryRewards(c *gin.Context) {
	id := c.Param("userId")
	today := time.Now().Truncate(24 * time.Hour)

	// DATE(received_at) returns string so we store strings here
	var dateStrings []string

	initializers.DB.
		Model(&models.RewardEvent{}).
		Select("DATE(received_at)").
		Where("user_id = ? AND received_at < ?", id, today).
		Group("DATE(received_at)").
		Order("DATE(received_at) DESC").
		Pluck("date", &dateStrings)

	history := make([]gin.H, 0)

	for _, ds := range dateStrings {
		// convert "2025-11-12" -> time.Time
		dayStart, _ := time.Parse("2006-01-02", ds)
		dayEnd := dayStart.Add(24 * time.Hour)

		var rewards []models.RewardEvent
		initializers.DB.
			Where("user_id = ? AND received_at >= ? AND received_at < ?", id, dayStart, dayEnd).
			Find(&rewards)

		dayTotal := 0.0
		for _, r := range rewards {
			price := generateRandomPrice(r.StockSymbol)
			dayTotal += price * r.Shares
		}

		dayTotal = float64(int(dayTotal*100)) / 100

		history = append(history, gin.H{
			"date":     ds,
			"inrValue": dayTotal,
		})
	}

	c.JSON(200, gin.H{"history": history})
}
