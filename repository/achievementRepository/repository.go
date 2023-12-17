package achievementRepository

import (
	"backend_scrolling_simulator/models"
	"backend_scrolling_simulator/repository"
	"gorm.io/gorm"
)

func UnlockAchievement(UserID uint, achievementId uint) error {
	db := repository.GetDB()
	return db.Model(
		models.User{Model: gorm.Model{ID: UserID}},
	).Association("Achievements").Append(
		&models.Achievement{Model: gorm.Model{ID: achievementId}},
	)
}

func GetAllAchievements() ([]models.Achievement, error) {
	db := repository.GetDB()
	var achievements []models.Achievement
	err := db.Find(&achievements).Error
	return achievements, err
}
