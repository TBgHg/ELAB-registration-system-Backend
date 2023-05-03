package thread

import (
	"elab-backend/internal/model/space/content"
	"elab-backend/internal/server/http/space/thread/comment"
	"elab-backend/internal/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ApplyGroup(engine *gin.RouterGroup) {
	group := engine.Group("/thread")
	{
		group.GET("", middlewares.SpacePublicPermissionMiddleware(), searchThread)
		group.POST("", middlewares.SpacePrivatePermissionMiddleware(), createThread)
		threadIdSpecifiedGroup := group.Group("/:thread_id")
		comment.ApplyGroup(threadIdSpecifiedGroup)
		threadIdSpecifiedGroup.Use(middlewares.ThreadIdRequiredMiddleware())
		threadIdSpecifiedGroup.PATCH("",
			middlewares.ContentOperatorMiddleware(content.Thread),
			middlewares.SpacePrivatePermissionMiddleware(),
			updateThreadById)
		threadIdSpecifiedGroup.DELETE("",
			middlewares.ContentOperatorMiddleware(content.Thread),
			middlewares.SpacePrivatePermissionMiddleware(),
			deleteThreadById)
		threadIdSpecifiedGroup.GET("/history/:history_id", middlewares.SpacePublicPermissionMiddleware(), getThreadHistoryById)
	}
}
