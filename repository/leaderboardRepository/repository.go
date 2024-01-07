package leaderboardRepository

import (
	"backend_scrolling_simulator/models"
	"backend_scrolling_simulator/repository"
	"fmt"
)

func GetLeaderboard(predicate models.LeaderboardTypePredicate) (*models.Leaderboard, error) {
	db := repository.GetDB()

	var leaderboardEntries []models.LeaderboardEntry
	err := db.Joins("JOIN users ON users.id = leaderboard.user_id").Where(&models.LeaderboardEntry{Timeframe: predicate.Timeframe}).Order(predicate.LeaderboardType + " DESC").Find(&leaderboardEntries).Error

	for i, entry := range leaderboardEntries {
		fmt.Printf("Entry %d: %v\n", i, entry.User)
	}

	leaderboard := models.NewLeaderboard(leaderboardEntries, predicate)
	return leaderboard, err
}
