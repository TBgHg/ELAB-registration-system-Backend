package comment

import (
	"elab-backend/internal/model/space/content"
	"elab-backend/internal/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ApplyGroup(engine *gin.RouterGroup) {
	group := engine.Group("/comment")
	{
		group.GET("", middlewares.SpacePublicPermissionMiddleware(), getComment)
		group.POST("", middlewares.SpacePrivatePermissionMiddleware(), createComment)
		commentIdSpecifiedGroup := group.Group("/:comment_id")
		commentIdSpecifiedGroup.Use(middlewares.CommentIdRequiredMiddleware())
		commentIdSpecifiedGroup.DELETE("",
			middlewares.ContentOperatorMiddleware(content.Comment),
			middlewares.SpacePrivatePermissionMiddleware(),
			deleteCommentById)
		commentIdSpecifiedGroup.POST("",
			middlewares.ContentOperatorMiddleware(content.Comment),
			middlewares.SpacePrivatePermissionMiddleware(),
			updateCommentById)
		commentLikeGroup := commentIdSpecifiedGroup.Group("/:comment_id/like")
		commentLikeGroup.POST("", middlewares.SpacePrivatePermissionMiddleware(), likeCommentById)
		commentLikeGroup.DELETE("", middlewares.SpacePrivatePermissionMiddleware(), dislikeCommentById)
	}
}
