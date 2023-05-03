package middlewares

import (
	"elab-backend/internal/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func UserOpenidRequiredMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		openid := ctx.Param("openid")
		if openid == "" {
			slog.Error("abort with status 400", "error", "openid_required", "openid", openid)
			ctx.AbortWithStatusJSON(400, model.NewInvalidParamError())
		}
		ctx.Next()
	}
}
