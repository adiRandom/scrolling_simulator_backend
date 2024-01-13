package engagement

import (
	"backend_scrolling_simulator/dtos"
	"backend_scrolling_simulator/repository/contentRepository"
	"backend_scrolling_simulator/repository/likeRepository"
	"backend_scrolling_simulator/services/post"
	"fmt"
)

func LikePost(postId uint, userId uint) (dtos.ReactResponse, *dtos.Error) {
	postExists, err := contentRepository.PostExists(postId)
	if err != nil {
		return dtos.ReactResponse{}, &dtos.Error{
			Msg:  fmt.Sprintf("Error with post id %d", postId),
			Code: 500,
		}
	}

	if !postExists {
		return dtos.ReactResponse{}, &dtos.Error{
			Msg:  fmt.Sprintf("Post with id %d does not exist", postId),
			Code: 404,
		}
	}

	err = likeRepository.LikePost(postId, userId)
	if err != nil {
		return dtos.ReactResponse{}, &dtos.Error{
			Msg:  fmt.Sprintf("Error liking post with id %d", postId),
			Code: 500,
		}
	}

	reaction, err := post.GetReaction(postId)
	if err != nil {
		return dtos.ReactResponse{}, &dtos.Error{
			Msg:  fmt.Sprintf("Error getting reaction for post with id %d", postId),
			Code: 500,
		}
	}

	return reaction.ToDto(), nil
}

func DislikePost(postId uint, userId uint) (dtos.ReactResponse, *dtos.Error) {
	postExists, err := contentRepository.PostExists(postId)
	if err != nil {
		return dtos.ReactResponse{}, &dtos.Error{
			Msg:  fmt.Sprintf("Error with post id %d", postId),
			Code: 500,
		}
	}

	if !postExists {
		return dtos.ReactResponse{}, &dtos.Error{
			Msg:  fmt.Sprintf("Post with id %d does not exist", postId),
			Code: 404,
		}
	}

	err = likeRepository.DislikePost(postId, userId)
	if err != nil {
		return dtos.ReactResponse{}, &dtos.Error{
			Msg:  fmt.Sprintf("Error disliking post with id %d", postId),
			Code: 500,
		}
	}

	reaction, err := post.GetReaction(postId)
	if err != nil {
		return dtos.ReactResponse{}, &dtos.Error{
			Msg:  fmt.Sprintf("Error getting reaction for post with id %d", postId),
			Code: 500,
		}
	}

	return reaction.ToDto(), nil
}
