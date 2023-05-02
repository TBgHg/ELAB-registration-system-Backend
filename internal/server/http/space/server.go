package space

import (
	"elab-backend/internal/server/http/space/member"
	"elab-backend/internal/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ApplyGroup(engine *gin.RouterGroup) {
	group := engine.Group("/space")
	{
		group.GET("", searchSpace)
		group.POST("", createSpace)
		idSpecifiedGroup := group.Group("/:space_id")
		idSpecifiedGroup.Use(middlewares.SpaceIdRequiredMiddleware())
		{
			idSpecifiedGroup.GET("", middlewares.SpacePublicPermissionMiddleware(), getSpaceById)
			idSpecifiedGroup.DELETE("", deleteSpaceById)
		}
		member.ApplyGroup(group)
	}
}
