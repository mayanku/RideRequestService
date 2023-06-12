package initializers

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() (*gorm.DB, error) {
	var err error
	dsn := os.Getenv("DB")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if DB != nil {
		fmt.Print("connected to db")
		return DB, nil
	}

	if err != nil {
		err_mod := errors.New("Can't open database: " + err.Error())
		return DB, err_mod
	}
	return DB, nil

}
