package models

import (
	"gorm.io/gorm"
)

type Achievement struct {
	gorm.Model
	Users        []*User `gorm:"many2many:user_achievements;"`
	Name         string  `gorm:"not null"`
	Description  string  `gorm:"not null"`
	BronzePoints uint    `gorm:"not null"`
	SilverPoints uint    `gorm:"not null"`
	GoldPoints   uint    `gorm:"not null"`
	IconUrl      string  `gorm:"not null"`
}
