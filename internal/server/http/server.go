package http

import (
	"elab-backend/internal/service"

	"github.com/gin-gonic/gin"
)

var svc *service.Service

func Init(r *gin.Engine, s *service.Service) {
	svc = s
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })

	user := r.Group("/user")
	{
		user.Use(JwtMiddleware())
		user.GET("/:openid", getUser)
		// todo:没有进行本地测试，不知道这个接口是否能用
		user.PUT("/:openid/avatar", uploadAvatar)
		// todo:待space表创建完成后补充
		user.GET("/:openid/space", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })
	}

	a := r.Group("/application")
	{
		a.Use(JwtMiddleware())
		// 获取用户申请状态信息，若科中成员则返回错误
		// 主要返回 是否提交了申请表，是否选定了面试场次
		// 若选定了面试场次，则一并返回面试场次信息，和面试签到码的内容(其实就是面试场次+openid后附带私钥签名，麻烦就直接用base64
		//a.GET("/", getApplication)
		// 获取申请表单
		//a.GET("/form")
	}

	oldUser := r.Group("/user")
	{
		oldUser.Use(JwtMiddleware())
		// 提交报名信息
		oldUser.POST("/submit", userSubmit)
		// 获取报名信息
		//oldUser.GET("/get", getUser)
		// 修改报名信息
		oldUser.POST("/update", userUpdate)
	}
	interview := r.Group("/interview")
	{
		interview.Use(JwtMiddleware())
		// 选择面试场次，新场次id为0表示取消面试
		interview.POST("/select", interviewSelect)
		// 获取面试信息-如果已选择面试场次，显示个人场次信息及结果，如果未选择面试场次，显示所有场次信息
		interview.GET("/get", interviewGet)
	}
	auth := r.Group("/auth")
	{
		auth.GET("/callback", callback)
	}
}
