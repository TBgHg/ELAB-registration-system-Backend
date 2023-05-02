package member

import "github.com/gin-gonic/gin"

func ApplyGroup(engine *gin.RouterGroup) {
	group := engine.Group("/member")
	{
		group.GET("", getMemberList)
		group.POST("", operateMember)
		group.POST("/:openid", operateMemberById)
		group.DELETE("/:openid", deleteMemberById)
	}
}
