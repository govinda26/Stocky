package main

import (
	"stocky/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
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

	r.Run()
}
