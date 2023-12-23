package dtos

type LeaderboardType struct {
	LeaderboardType string `form:"leaderboard_type" binding:"required"`
	Timeframe       string `form:"timeframe" binding:"required"`
}

type LeaderboardEntry struct {
	User  User `json:"user"`
	Value uint `json:"value"`
}
