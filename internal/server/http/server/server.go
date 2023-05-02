package server

import (
	"elab-backend/internal/server/http/application"
	"elab-backend/internal/server/http/auth"
	"elab-backend/internal/server/http/space"
	"elab-backend/internal/server/http/user"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	group := engine.Group("/v1")
	engine.Use(gin.Logger(), gin.Recovery(), requestid.New())
	auth.ApplyGroup(group)
	user.ApplyGroup(group)
	application.ApplyGroup(group)
	space.ApplyGroup(group)
}
