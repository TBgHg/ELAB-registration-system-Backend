package comment

import "github.com/gin-gonic/gin"

func ApplyGroup(engine *gin.RouterGroup) {
	group := engine.Group("/comment")
	{
		group.GET("", searchComment)
		group.POST("", newComment)
		group.GET("/:comment_id", getCommentByID)
		group.PATCH("/:comment_id", updateCommentByID)
		group.DELETE("/:comment_id", deleteCommentByID)
		group.POST("/:comment_id/like", likeCommentByID)
		group.DELETE("/:comment_id/like", unlikeCommentByID)
	}
}
