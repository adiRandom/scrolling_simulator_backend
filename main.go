package main

import (
	"backend_scrolling_simulator/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	engine := gin.Default()
	loadRoutes(engine)
	err = engine.Run(":" + os.Getenv("SCROLLING_SIMULATOR_API_PORT"))
	if err != nil {
		return
	}
}

func loadRoutes(e *gin.Engine) {
	routes.LoadEngagementGroup(e)
	routes.LoadAchievementRoutes(e)
}
