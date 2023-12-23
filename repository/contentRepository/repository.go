package contentRepository

import (
	"backend_scrolling_simulator/models"
	"backend_scrolling_simulator/repository"
	"fmt"
)

func GetRandomPostsUnseenBy(userId uint, count int) ([]models.Post, error) {
	var posts []models.Post
	db := repository.GetDB()
	err := db.Model(&models.Post{}).
		Joins(fmt.Sprintf(
			"LEFT JOIN seen_posts ON posts.id = seen_posts.post_id AND seen_posts.user_id = %d ",
			userId),
		).
		Where("seen_posts.post_id IS NULL").
		Order("RANDOM()").
		Limit(count).
		Preload("Topics").
		Preload("Tags").
		Find(&posts).Error
	return posts, err
}
