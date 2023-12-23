package models

import (
	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Emoji string `gorm:"not null"`
}