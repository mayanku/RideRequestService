package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"riderequest/initializers"
	"riderequest/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	var err error
	dsn := os.Getenv("DB")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.SlotTable{})

}
