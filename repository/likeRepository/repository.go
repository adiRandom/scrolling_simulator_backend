package likeRepository

import (
	"backend_scrolling_simulator/models"
	"backend_scrolling_simulator/repository"
)

func LikePost(postId uint, UserID uint) error {
	db := repository.GetDB()
	db.Create(&models.Like{UserID: UserID, PostId: postId, IsLike: true})
	return db.Where("post_id = ? AND user_id = ? AND is_like = false", postId, UserID).Delete(&models.Like{}).Error
}

func DislikePost(postId uint, UserID uint) error {
	db := repository.GetDB()
	db.Create(&models.Like{UserID: UserID, PostId: postId, IsLike: false})
	return db.Where("post_id = ? AND user_id = ? AND is_like = true", postId, UserID).Delete(&models.Like{}).Error
}
