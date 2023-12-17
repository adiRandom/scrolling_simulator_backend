package routes

import (
	"backend_scrolling_simulator/dtos"
	"backend_scrolling_simulator/lib"
	"backend_scrolling_simulator/lib/routes"
	"backend_scrolling_simulator/repository/leaderboardRepository"
	"github.com/gin-gonic/gin"
)

func GetLeaderboard(ctx *gin.Context) {
	ctxWrapper, err := routes.GetContextWrapper[dtos.LeaderboardType, lib.None](ctx)
	if err != nil {
		return
	}

	getLeaderboardDto := ctxWrapper.GetQueryParams()

	leaderboard, err := leaderboardRepository.GetLeaderboard(getLeaderboardDto.ToPredicate())
	if err != nil {
		ctxWrapper.ReturnErrorResponse(err, 500)
		return
	}

	ctx.JSON(200, leaderboard)
}

func LoadLeaderboardRoutes(e *gin.Engine) {
	e.GET("/leaderboard", GetLeaderboard)
}
