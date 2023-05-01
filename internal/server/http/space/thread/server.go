package thread

import "github.com/gin-gonic/gin"

func ApplyGroup(engine *gin.RouterGroup) {
	group := engine.Group("/thread")
	{
		group.GET("", searchThread)
		group.POST("", newThread)
		group.GET("/:thread_id", getThreadByID)
		group.DELETE("/:thread_id", deleteThreadByID)
	}
}
