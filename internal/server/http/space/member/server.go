package member

import "github.com/gin-gonic/gin"

func ApplyGroup(engine *gin.RouterGroup) {
	group := engine.Group("/member")
	{
		group.GET("", searchMember)
		group.POST("", newMember)
		group.POST("/:member_id", updateMemberByID)
		group.DELETE("/:member_id", deleteMemberByID)
	}
}
