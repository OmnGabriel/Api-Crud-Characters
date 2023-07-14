package database

import (
	"log"

	"github.com/OmnGabriel/go-api-rest.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectWithDatabase() {
	connectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("Error connecting to database")
	}
	DB.AutoMigrate(&models.Character{})
}
