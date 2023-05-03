package space

import (
	"elab-backend/internal/server/http/space/member"
	"elab-backend/internal/server/http/space/thread"
	"elab-backend/internal/server/http/space/wiki"
	"elab-backend/internal/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ApplyGroup(engine *gin.RouterGroup) {
	group := engine.Group("/space")
	group.Use(middlewares.OAuthLoginRequiredMiddleware())
	{
		group.GET("", searchSpace)
		group.POST("", createSpace)
		spaceIdSpecifiedGroup := group.Group("/:space_id")
		spaceIdSpecifiedGroup.Use(middlewares.SpaceIdRequiredMiddleware())
		spaceIdSpecifiedGroup.GET("/:space_id",
			middlewares.SpacePublicPermissionMiddleware(),
			getSpaceById,
		)
		spaceIdSpecifiedGroup.DELETE("/:space_id", deleteSpaceById)
		member.ApplyGroup(spaceIdSpecifiedGroup)
		thread.ApplyGroup(spaceIdSpecifiedGroup)
		wiki.ApplyGroup(spaceIdSpecifiedGroup)
	}
}
