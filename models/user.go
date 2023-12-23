package models

import (
	"backend_scrolling_simulator/dtos"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Achievements []*Achievement `gorm:"many2many:user_achievements;"`
}

func (u *User) ToDto() dtos.User {
	return dtos.User{
		Id: u.ID,
	}
}
