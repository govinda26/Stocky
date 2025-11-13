package controllers

import (
	"math"
	"math/rand"
	"stocky/initializers"
	"stocky/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func generateRandomPrice(symbol string) float64 {
	return 1000 + rand.Float64()*1000
}

func RewardCreate(c *gin.Context) {
	//STEP 1: Create request body struct
	//struct to hold request data
	var body struct {
		UserID      uint    `json:"userId"`
		StockSymbol string  `json:"stockSymbol"`
		Shares      float64 `json:"shares"`
	}

	//STEP 2: Bind JSON to struct
	//adds request data into above struct
	c.Bind(&body)

	//STEP 3: Validate input
	if body.UserID == 0 || body.StockSymbol == "" || body.Shares <= 0 {
		c.JSON(400, gin.H{
			"error": "Incorrect user input",
		})
		return
	}

	//STEP 4: Insert RewardEvent
	//create a reward struct
	reward := models.RewardEvent{
		UserID:      body.UserID,
		StockSymbol: body.StockSymbol,
		Shares:      body.Shares,
		ReceivedAt:  time.Now(),
		ExternalRef: uuid.NewString(),
	}

	//pass the struct
	result := initializers.DB.Create(&reward)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"Error": "Something went wrong while inserting data into the table",
		})
		return
	}

	//STEP 5: Update Holdings
	//initialise struct of holding
	var holding models.Holding

	//because holding is of type models.Holding DB knows where to look
	initializers.DB.Where("user_id = ? AND stock_symbol = ?", body.UserID, body.StockSymbol).First(&holding)

	//.First() fills holding with id found from above query
	if holding.ID == 0 {
		//no holding exists
		holding = models.Holding{
			UserID:      body.UserID,
			StockSymbol: body.StockSymbol,
			TotalShares: body.Shares,
		}
		initializers.DB.Create(&holding)
	} else {
		//holding exists
		holding.TotalShares = holding.TotalShares + body.Shares
		initializers.DB.Save(&holding)
	}

	// Generate Random Price
	price := generateRandomPrice(body.StockSymbol)
	//total Share price
	totalCost := price * body.Shares
	totalCost = math.Round(totalCost*100) / 100
	//brokerage
	fee := totalCost * 0.02
	fee = math.Round(fee*100) / 100
	//for ledger entry
	groupId := uuid.NewString()

	//Add Ledger Entries
	//company gets shares
	ledger1 := models.Ledger{
		EntryGroupId: groupId,
		Account:      "stock:" + body.StockSymbol,
		Debit:        totalCost,
		Credit:       0,
		Shares:       body.Shares,
		StockSymbol:  body.StockSymbol,
	}
	initializers.DB.Create(&ledger1)

	//company pays cash to NSE / BSE for shares
	ledger2 := models.Ledger{
		EntryGroupId: groupId,
		Account:      "cash",
		Debit:        0,
		Credit:       totalCost,
	}
	initializers.DB.Create(&ledger2)

	//company pays fees (brokerage)
	ledger3 := models.Ledger{
		EntryGroupId: groupId,
		Account:      "fees",
		Debit:        0,
		Credit:       fee,
	}
	initializers.DB.Create(&ledger3)

	// Return response
	c.JSON(200, gin.H{
		"reward": reward,
	})
}
