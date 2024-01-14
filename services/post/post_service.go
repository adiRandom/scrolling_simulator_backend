package post

import (
	"backend_scrolling_simulator/models"
	"backend_scrolling_simulator/repository/likeRepository"
)

func clampLikeDislikeRatio(postId uint) (float64, error) {
	ratio, err := likeRepository.GetLikeDislikeRatio(postId)
	if err != nil {
		return 0, err
	}

	switch {
	case ratio < 0.45:
		return 0.3, nil

	case ratio < 0.6:
		return 0.5, nil

	case ratio < 0.9:
		return 0.6, nil

	case ratio < 1.3:
		return 1, nil

	case ratio < 2:
		return 1.5, nil

	case ratio < 3:
		return 2, nil

	case ratio > 3:
		return 3, nil

	default:
		return 1, nil
	}
}

func GetReaction(postId uint) (models.ReactText, error) {
	likeDislikeRatio, err := clampLikeDislikeRatio(postId)
	if err != nil {
		return models.ReactText{}, err
	}

	return likeRepository.GetReaction(postId, likeDislikeRatio)
}
