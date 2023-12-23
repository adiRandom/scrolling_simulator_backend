package models

import "gorm.io/gorm"

type LeaderboardEntry struct {
	gorm.Model
	UserID    uint   `gorm:"not null"`
	Points    uint   `gorm:"not null"`
	Timeframe string `gorm:"not null"`
	Distance  uint   `gorm:"not null"`
	User      User
}

func (_ LeaderboardEntry) TableName() string {
	return "leaderboard"
}
