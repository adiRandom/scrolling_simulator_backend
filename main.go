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
	println(os.Getenv("SCROLLING_SIMULATOR_API_PORT"))

	engine := gin.Default()
	loadRoutes(engine)
}

func loadRoutes(e *gin.Engine) {
	routes.LoadEngagementGroup(e)
}
