package models

import "gorm.io/gorm"

type CrossPost struct {
	gorm.Model
	UserID uint `gorm:"not null"`
	PostId uint `gorm:"not null"`
}
