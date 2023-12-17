package routes

import (
	"backend_scrolling_simulator/dtos"
	"backend_scrolling_simulator/lib"
	"backend_scrolling_simulator/lib/routes"
	"backend_scrolling_simulator/repository/crossPostRepository"
	"backend_scrolling_simulator/repository/likeRepository"
	"github.com/gin-gonic/gin"
)

func like(ctx *gin.Context) {
	ctxWrapper, err := routes.GetContextWrapper[lib.None, dtos.Like](ctx)
	if err != nil {
		return
	}

	likeDto := ctxWrapper.GetBody()
	user := ctxWrapper.GetCurrentUser()

	err = likeRepository.LikePost(likeDto.PostId, user.ID)
	if err != nil {
		ctxWrapper.ReturnErrorResponse(err, 500)
		return
	}

	ctx.JSON(204, nil)
}

func dislike(ctx *gin.Context) {
	ctxWrapper, err := routes.GetContextWrapper[lib.None, dtos.Like](ctx)
	if err != nil {
		return
	}

	dislikeDto := ctxWrapper.GetBody()
	user := ctxWrapper.GetCurrentUser()

	err = likeRepository.DislikePost(dislikeDto.PostId, user.ID)
	if err != nil {
		ctxWrapper.ReturnErrorResponse(err, 500)
		return
	}

	ctx.JSON(204, nil)
}

func crossPost(ctx *gin.Context) {
	ctxWrapper, err := routes.GetContextWrapper[lib.None, dtos.CrossPost](ctx)
	if err != nil {
		return
	}

	crossPostDto := ctxWrapper.GetBody()
	user := ctxWrapper.GetCurrentUser()

	err = crossPostRepository.CreateCrossPost(crossPostDto.PostId, user.ID)
	if err != nil {
		ctxWrapper.ReturnErrorResponse(err, 500)
		return
	}

	ctx.JSON(204, nil)
}

func LoadEngagementGroup(e *gin.Engine) {
	g := e.Group("/engagement")

	g.POST("/like", like)
	g.POST("/dislike", dislike)
	g.POST("/crosspost", crossPost)
}
