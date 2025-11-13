package controllers

import (
	"stocky/initializers"
	"stocky/models"
	"time"

	"github.com/gin-gonic/gin"
)

func TodayStocks(c *gin.Context) {
	//get userId from params
	id := c.Param("userId")

	//get today's date
	today := time.Now().Truncate(24 * time.Hour)

	//get rewards
	//because we are expecting multiple rewards we will slice to store them
	var rewards []models.RewardEvent

	initializers.DB.Where("user_id = ? AND received_at >= ?", id, today).Find(&rewards)

	//return
	c.JSON(200, gin.H{
		"todayRewards": rewards,
	})
}
