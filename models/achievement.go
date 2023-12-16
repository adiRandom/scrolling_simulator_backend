package models

import (
	"backend_scrolling_simulator/dtos"
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

func (achievement *Achievement) ToDto(isUnlocked bool) dtos.Achievement {
	return dtos.Achievement{
		Id:           achievement.ID,
		Name:         achievement.Name,
		Description:  achievement.Description,
		BronzePoints: achievement.BronzePoints,
		SilverPoints: achievement.SilverPoints,
		GoldPoints:   achievement.GoldPoints,
		IconUrl:      achievement.IconUrl,
		Unlocked:     isUnlocked,
	}
}
