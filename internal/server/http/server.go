package http

import (
	"ELAB-registration-system-Backend/internal/service"
	"github.com/gin-gonic/gin"
)

var svc *service.Service

func Init(r *gin.Engine, s *service.Service) {
	svc = s
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })

	// 添加中间件，验证token
	//r.GET("/callback", service.VerifyToken(), func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })
	r.Use(JwtMiddleware())
	signup := r.Group("/signup")
	{
		// 提交报名信息
		signup.POST("/submit", signupSubmit)
		// 获取报名信息
		signup.GET("/get", signupGet)
		// 修改报名信息
		signup.POST("/update", signupUpdate)
	}
	interview := r.Group("/interview")
	{
		// 选择面试场次，新场次id为0表示取消面试
		interview.POST("/select", interviewSelect)
		// 获取面试信息-如果已选择面试场次，显示个人场次信息及结果，如果未选择面试场次，显示所有场次信息
		interview.GET("/get", interviewGet)
	}
}
