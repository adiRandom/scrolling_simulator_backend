package user

import (
	"backend_scrolling_simulator/models"
	"gorm.io/gorm"
)

func GetCurrentUser(token string) (models.User, error) {
	return models.User{Model: gorm.Model{ID: 1}}, nil
}
