package application

import (
	"elab-backend/internal/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ApplyGroup(engine *gin.RouterGroup) {
	group := engine.Group("/application")
	group.Use(middlewares.OAuthLoginRequiredMiddleware())
	{
		group.GET("/status", getStatus)
		group.GET("/longTextForm", getLongTextForm)
		group.PATCH("/longTextForm", updateLongTextForm)
		group.GET("/room", getRoom)
		group.POST("/room", setRoom)
		group.DELETE("/room", deleteRoom)
	}
}
