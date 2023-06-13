package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"riderequest/initializers"
	"riderequest/models"
	booking_status "riderequest/utils/enums"
	"time"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func BookingRequest(c *gin.Context) {
	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	rideRequest := models.RideRequest{
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
		BookingID:            "12345",
		UserID:               "user123",
		BookingDate:          "2023-06-12",
		RideDate:             "2023-06-13",
		SourceAddress:        "123 Main St",
		SourceCity:           "CityA",
		SourceState:          "StateA",
		SourcePincode:        12345,
		SourceLongitude:      45.6789,
		SourceLatitude:       67.8901,
		DestinationLatitude:  12.3456,
		DestinationLongitude: 23.4567,
		DestinationAddress:   "456 Park Ave",
		DestinationCity:      "CityB",
		DestinationState:     "StateB",
		DestinationPincode:   54321,
		BookingTime:          "10:00 AM",
		RideTime:             time.Now(),
		SlotTime:             "Morning",
		SlotNumber:           1,
		Distance:             "10 miles",
		Amount:               123,
		Status:               booking_status.BookingStatus(0),
	}
	result := db.Create(&rideRequest)
	if result.Error != nil {
		// Handle error when the table does not return anything
		fmt.Println(result)
	}
}
