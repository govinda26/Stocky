package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// creating a global DB variable which holds value of gorm.DB
var DB *gorm.DB

func ConnectToDB() {
	//creating an err
	var err error

	//postgres Connection string from env
	dsn := os.Getenv("DB_URL")

	//establishing connection with postgres
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while establishing connection to db, ", err)
	}
}
