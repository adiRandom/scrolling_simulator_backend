package dtos

import "backend_scrolling_simulator/repository/leaderboardRepository"

type LeaderboardType struct {
	LeaderboardType string
	Timeframe       string
}

func (l *LeaderboardType) ToPredicate() leaderboardRepository.LeaderboardTypePredicate {
	return leaderboardRepository.LeaderboardTypePredicate{
		LeaderboardType: l.LeaderboardType,
		Timeframe:       l.Timeframe,
	}
}
