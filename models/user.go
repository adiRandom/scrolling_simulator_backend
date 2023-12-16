package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Achievements []*Achievement `gorm:"many2many:user_achievements;"`
}
