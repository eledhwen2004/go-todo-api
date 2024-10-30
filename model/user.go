package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"AutoCreateTime"`
	UpdatedAt time.Time `gorm:"AutoUpdateTime"`
}
