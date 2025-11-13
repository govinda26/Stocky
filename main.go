package main

import (
	"stocky/controllers"
	"stocky/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.InitLogger()
}

func main() {
	//create a gin router
	r := gin.Default()

	//Define a simple GET endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Stocks": "Working",
		})
	})

	r.POST("/reward", controllers.RewardCreate)
	r.GET("/today-stocks/:userId", controllers.TodayStocks)
	r.GET("/stats/:userId", controllers.Stats)
	r.GET("/historical-inr/:userId", controllers.HistoryRewards)

	r.Run()
}
