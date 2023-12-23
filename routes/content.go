package routes

import (
	"backend_scrolling_simulator/dtos"
	"backend_scrolling_simulator/lib"
	"backend_scrolling_simulator/lib/functional"
	"backend_scrolling_simulator/lib/routes"
	"backend_scrolling_simulator/models"
	"backend_scrolling_simulator/repository/contentRepository"
	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context) {
	ctxWrapper, err := routes.GetContextWrapper[dtos.Pagination, lib.None](ctx)
	if err != nil {
		return
	}

	pagination := ctxWrapper.GetQueryParams()
	user := ctxWrapper.GetCurrentUser()
	posts, err := contentRepository.GetRandomPostsUnseenBy(user.ID, pagination.Limit)
	if err != nil {
		ctxWrapper.ReturnErrorResponse(err, 500)
		return
	}

	response := functional.Map(posts, func(post models.Post) dtos.Post { return dtos.NewPostFromModel(post) })

	ctx.JSON(200, response)
}

func LoadPostRoutes(engine *gin.Engine) {
	postGroup := engine.Group("/posts")
	{
		postGroup.GET("", GetPosts)
	}
}
