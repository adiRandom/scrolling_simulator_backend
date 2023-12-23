package models

import (
	"backend_scrolling_simulator/dtos"
	"backend_scrolling_simulator/lib"
	"backend_scrolling_simulator/lib/functional"
	"gorm.io/gorm"
)

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

const DAILY_LEADERBOARD = "daily"
const WEEKLY_LEADERBOARD = "weekly"
const MONTHLY_LEADERBOARD = "monthly"

const POINTS_LEADERBOARD = "points"
const DISTANCE_LEADERBOARD = "distance"

type LeaderboardTypePredicate struct {
	LeaderboardType string
	Timeframe       string
}

func NewLeaderboardTypePredicate(leaderboardType string, timeframe string) *LeaderboardTypePredicate {
	return &LeaderboardTypePredicate{LeaderboardType: leaderboardType, Timeframe: timeframe}
}

type Leaderboard struct {
	entries   []LeaderboardEntry
	predicate LeaderboardTypePredicate
}

func NewLeaderboard(entries []LeaderboardEntry, predicate LeaderboardTypePredicate) *Leaderboard {
	return &Leaderboard{entries: entries, predicate: predicate}
}

func (l *Leaderboard) getValues() []uint {
	var values []uint
	for _, entry := range l.entries {
		switch l.predicate.LeaderboardType {
		case POINTS_LEADERBOARD:
			{
				values = append(values, entry.Points)
			}
		case DISTANCE_LEADERBOARD:
			{
				values = append(values, entry.Distance)
			}
		}
	}
	return values
}

func (l *Leaderboard) GetUsers() []User {
	return functional.Map(l.entries, func(entry LeaderboardEntry) User {
		return entry.User
	})
}

func (l *Leaderboard) ToDto() []dtos.LeaderboardEntry {
	values := l.getValues()
	zipped := functional.Zip(l.entries, values)
	mapped := functional.Map(zipped, func(pair lib.Pair[LeaderboardEntry, uint]) dtos.LeaderboardEntry {
		return dtos.LeaderboardEntry{User: pair.First.User.ToDto(), Value: pair.Second}
	})

	return mapped
}
