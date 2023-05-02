package auth

import (
	"github.com/gin-gonic/gin"
)

func ApplyGroup(engine *gin.RouterGroup) {
	group := engine.Group("/auth")
	{
		group.GET("/callback", callback)
		group.POST("/create", createLoginSession)
		group.POST("/refresh", refresh)
	}
}
