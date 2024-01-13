package likeRepository

import (
	"backend_scrolling_simulator/models"
	"backend_scrolling_simulator/repository"
)

func LikePost(postId uint, UserID uint) error {
	db := repository.GetDB()
	db.Create(&models.Like{UserID: UserID, PostId: postId, IsLike: true})
	return db.Where("post_id = ? AND user_id = ? AND is_like = false", postId, UserID).
		Delete(&models.Like{}).
		Error
}

func DislikePost(postId uint, UserID uint) error {
	db := repository.GetDB()
	db.Create(&models.Like{UserID: UserID, PostId: postId, IsLike: false})
	return db.Where("post_id = ? AND user_id = ? AND is_like = true", postId, UserID).
		Delete(&models.Like{}).
		Error
}

func GetLikeDislikeRatio(postId uint) (float64, error) {
	db := repository.GetDB()
	var ratio float64
	err := db.Model(&models.Like{}).
		Select("SUM(CASE WHEN is_like THEN 1 ELSE 0 END) / COALESCE(NULLIF(SUM(CASE WHEN is_like THEN 0 ELSE 1 END), 0), 1)").
		Where("post_id = ?", postId).
		Scan(&ratio).
		Error
	return ratio, err
}

func GetReaction(postId uint, ratio float64) (models.React, error) {
	db := repository.GetDB()
	var react models.React
	err := db.Model(&react).
		Select("react_text").
		Where("post_id = ? AND ratio = ?",
			postId,
			ratio,
		).
		Order("RANDOM()").
		Limit(1).
		Find(&react).
		Error
	return react, err
}
