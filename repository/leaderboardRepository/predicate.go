package leaderboardRepository

const DAILY_LEADERBOARD = "daily"
const WEEKLY_LEADERBOARD = "weekly"
const MONTHLY_LEADERBOARD = "monthly"

const POINTS_LEADERBOARD = "points"
const DISTANCE_LEADERBOARD = "distance"

type LeaderboardTypePredicate struct {
	LeaderboardType string
	Timeframe       string
}
