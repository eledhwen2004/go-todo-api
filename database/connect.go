package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		// Logger will be added
		fmt.Printf("Error while connecting database : %v\n", err)
	}
	DB = db
}
