package routes

import (
	"backend_scrolling_simulator/dtos"
	"backend_scrolling_simulator/lib"
	"backend_scrolling_simulator/lib/routes"
	"backend_scrolling_simulator/models"
	"backend_scrolling_simulator/repository/leaderboardRepository"
	"github.com/gin-gonic/gin"
)

func GetLeaderboard(ctx *gin.Context) {
	ctxWrapper, err := routes.GetCtxWithQuery[dtos.LeaderboardType](ctx)
	if err != nil {
		return
	}

	getLeaderboardDto := ctxWrapper.GetQueryParams()

	leaderboard, err := leaderboardRepository.GetLeaderboard(
		*models.NewLeaderboardTypePredicate(
			getLeaderboardDto.LeaderboardType,
			getLeaderboardDto.Timeframe,
		),
	)
	if err != nil {
		println(err)
		ctxWrapper.ReturnErrorResponse(lib.Error{Msg: "Something went wrong!"}, 500)
		return
	}

	ctxWrapper.ReturnJSON(200, leaderboard.ToDto())
}

func LoadLeaderboardRoutes(e *gin.Engine) {
	e.GET(BasePath+"/leaderboard", GetLeaderboard)
}
