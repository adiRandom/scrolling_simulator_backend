package crossPostRepository

import (
	"backend_scrolling_simulator/models"
	"backend_scrolling_simulator/repository"
)

func CreateCrossPost(UserID uint, postId uint) error {
	db := repository.GetDB()
	return db.Create(&models.CrossPost{UserID: UserID, PostId: postId}).Error
}
