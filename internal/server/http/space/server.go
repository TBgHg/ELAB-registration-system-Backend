package space

import "github.com/gin-gonic/gin"

func ApplyGroup(engine *gin.RouterGroup) {
	group := engine.Group("/space")
	{
		group.GET("", searchSpace)
		group.POST("", newSpace)
		group.GET("/:space_id", getSpaceByID)
		group.DELETE("/:space_id", deleteSpaceByID)
	}
}
