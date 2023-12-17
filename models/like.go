package models

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserID uint `gorm:"not null"`
	PostId uint `gorm:"not null"`
	IsLike bool `gorm:"not null"`
}
