package crossPostRepository

import (
	"backend_scrolling_simulator/models"
	"backend_scrolling_simulator/repository"
)

func CreateCrossPost(userId uint, postId uint) error {
	db := repository.GetDB()
	return db.Create(&models.CrossPost{UserId: userId, PostId: postId}).Error
}
