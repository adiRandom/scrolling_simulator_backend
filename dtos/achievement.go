package dtos

import "backend_scrolling_simulator/models"

type Achievement struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	BronzePoints uint   `json:"bronzePoints"`
	SilverPoints uint   `json:"silverPoints"`
	GoldPoints   uint   `json:"goldPoints"`
	IconUrl      string `json:"iconUrl"`
	Unlocked     bool   `json:"unlocked"`
}

func NewAchievementFromModel(achievement models.Achievement, isUnlocked bool) Achievement {
	return Achievement{
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
