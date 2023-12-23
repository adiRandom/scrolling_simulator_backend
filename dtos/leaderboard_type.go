package dtos

import "backend_scrolling_simulator/repository/leaderboardRepository"

type LeaderboardType struct {
	LeaderboardType string `form:"leaderboard_type" binding:"required"`
	Timeframe       string `form:"timeframe" binding:"required"`
}

func (l *LeaderboardType) ToPredicate() leaderboardRepository.LeaderboardTypePredicate {
	return leaderboardRepository.LeaderboardTypePredicate{
		LeaderboardType: l.LeaderboardType,
		Timeframe:       l.Timeframe,
	}
}
