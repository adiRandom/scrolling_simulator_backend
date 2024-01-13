package routes

import (
	"backend_scrolling_simulator/dtos"
	"backend_scrolling_simulator/lib/routes"
	"backend_scrolling_simulator/repository/crossPostRepository"
	"backend_scrolling_simulator/services/engagement"
	"github.com/gin-gonic/gin"
)

func like(ctx *gin.Context) {
	ctxWrapper, err := routes.GetCtxWithPath[dtos.PostIdPathParam](ctx)
	if err != nil {
		return
	}

	pathParams := ctxWrapper.GetPathParams()
	user := ctxWrapper.GetCurrentUser()

	response, apiErr := engagement.LikePost(pathParams.PostId, user.ID)
	if err != nil {
		ctxWrapper.ReturnErrorResponse(apiErr, apiErr.Code)
		return
	}

	ctxWrapper.ReturnJSON(200, response)
}

func dislike(ctx *gin.Context) {
	ctxWrapper, err := routes.GetCtxWithPath[dtos.PostIdPathParam](ctx)
	if err != nil {
		return
	}

	pathParams := ctxWrapper.GetPathParams()

	user := ctxWrapper.GetCurrentUser()

	response, apiErr := engagement.DislikePost(pathParams.PostId, user.ID)
	if err != nil {
		ctxWrapper.ReturnErrorResponse(apiErr, apiErr.Code)
		return
	}

	ctxWrapper.ReturnJSON(200, response)
}

func crossPost(ctx *gin.Context) {
	ctxWrapper, err := routes.GetCtxWithPath[dtos.PostIdPathParam](ctx)
	if err != nil {
		return
	}

	pathParams := ctxWrapper.GetPathParams()

	user := ctxWrapper.GetCurrentUser()

	err = crossPostRepository.CreateCrossPost(pathParams.PostId, user.ID)
	if err != nil {
		ctxWrapper.ReturnErrorResponse(err, 500)
		return
	}

	ctxWrapper.ReturnJSON(204, nil)
}

func LoadEngagementGroup(e *gin.Engine) {
	g := e.Group(BasePath + "/engagement/:postId")

	g.POST("/like", like)
	g.POST("/dislike", dislike)
	g.POST("/crosspost", crossPost)
}
