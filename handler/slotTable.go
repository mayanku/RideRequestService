package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

// SlotTable The function for returning the slot table
func SlotTable(c *gin.Context) {
	var err error
	var slottable []models.SlotTable
	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	result := db.Find(&slottable)
	if result.Error != nil {
		// Handle error when the table does not return anything
	}

	fmt.Println(slottable) // for verifying the slot table output
	//fmt.Println(slottable)

}
