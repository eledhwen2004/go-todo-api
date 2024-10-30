package model

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title     string    `gorm:""`
	Content   string    `gorm:""`
	UserID    string    `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID;not null;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `gorm:"AutoCreateTime"`
	UpdatedAt time.Time `gorm:"AutoUpdateTime"`
}
