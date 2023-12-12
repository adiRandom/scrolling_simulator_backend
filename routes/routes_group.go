package routes

import "github.com/gin-gonic/gin"

type Group interface {
	LoadGroup(engine *gin.Engine)
}
