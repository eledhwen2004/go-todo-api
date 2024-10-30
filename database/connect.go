package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDatabase(dsn string) {
	DB, err = gorm.Open(postgres.Open(dsn))
}
