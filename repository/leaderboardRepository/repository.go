package leaderboardRepository

import (
	"backend_scrolling_simulator/models"
	"backend_scrolling_simulator/repository"
)

func GetLeaderboard(predicate LeaderboardTypePredicate) ([]*models.LeaderboardEntry, error) {
	db := repository.GetDB()

	var leaderboardEntries []*models.LeaderboardEntry
	err := db.Where("entry_type = ?",
		predicate.Timeframe).Order(predicate.LeaderboardType + " DESC").Find(&leaderboardEntries).Error

	return leaderboardEntries, err
}
