package initializers

import (
	"errors"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SyncDatabase() {

	if DB == nil {
		var err error
		dsn := os.Getenv("DB")

		// Open a db connection
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			errors.New("Can't open database: " + err.Error())
		}
	}

	/*if DB != nil {
		DB.AutoMigrate(&models.User{}, &models.Otp{})
	}*/
}
