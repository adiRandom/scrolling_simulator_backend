package leaderboardRepository

import (
	"backend_scrolling_simulator/models"
	"backend_scrolling_simulator/repository"
)

func GetLeaderboard(predicate models.LeaderboardTypePredicate) (*models.Leaderboard, error) {
	db := repository.GetDB()

	var leaderboardEntries []models.LeaderboardEntry
	err := db.Where(&models.LeaderboardEntry{Timeframe: predicate.Timeframe}).Order(predicate.LeaderboardType + " DESC").Joins("Users").Find(&leaderboardEntries).Error

	leaderboard := models.NewLeaderboard(leaderboardEntries, predicate)
	return leaderboard, err
}
