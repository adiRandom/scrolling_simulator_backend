package routes

import (
	"backend_scrolling_simulator/dtos"
	"backend_scrolling_simulator/lib"
	"backend_scrolling_simulator/lib/routes"
	"backend_scrolling_simulator/models"
	"backend_scrolling_simulator/repository/achievementRepository"
	"github.com/gin-gonic/gin"
)

func getAchievements(ctx *gin.Context) {
	ctxWrapper, err := routes.GetEmptyCtx(ctx)
	if err != nil {
		return
	}

	user := ctxWrapper.GetCurrentUser()

	achievements, err := achievementRepository.GetAllAchievements()
	if err != nil {
		ctxWrapper.ReturnErrorResponse(err, 500)
		return
	}

	response := make([]dtos.Achievement, 0)
	for _, achievement := range achievements {
		isUnlocked := lib.ContainsFunc(user.Achievements, func(el *models.Achievement) bool {
			return el.ID == achievement.ID
		})

		response = append(response, achievement.ToDto(isUnlocked))
	}

	ctxWrapper.ReturnJSON(200, response)
}

func LoadAchievementRoutes(router *gin.Engine) {
	router.GET(BasePath+"/achievements", getAchievements)
}
