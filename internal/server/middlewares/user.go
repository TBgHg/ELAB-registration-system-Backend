package middlewares

import (
	"elab-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func UserOpenidRequiredMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, exists := ctx.Get("openid")
		if !exists {
			ctx.AbortWithStatusJSON(400, model.NewInvalidParamError())
		}
		ctx.Next()
	}
}
