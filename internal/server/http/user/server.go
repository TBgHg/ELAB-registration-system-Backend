package user

import (
	"elab-backend/internal/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ApplyGroup(engine *gin.RouterGroup) {
	group := engine.Group("/user/:openid")
	group.Use(middlewares.UserOpenidRequiredMiddleware(), middlewares.OAuthLoginRequiredMiddleware())
	{
		group.GET("/", getUser)
		group.PATCH(
			"/",
			middlewares.OAuthSelfValidationMiddleware(),
			updateUser,
		)
		group.GET("/spaces", getUserSpaces)
	}
}
