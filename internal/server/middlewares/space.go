package middlewares

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func NewSpaceIdRequiredError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "space_id_required",
		Message: "/space/:space_id is required",
	}
}

func NewNoPermissionError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "no_permission",
		Message: "用户没有权限",
	}
}

func SpaceIdRequiredMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		spaceId := ctx.GetString("space_id")
		if spaceId == "" {
			ctx.AbortWithStatusJSON(400, NewSpaceIdRequiredError())
			return
		}
		ctx.Next()
	}
}

func SpacePublicPermissionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		spaceId := ctx.GetString("space_id")
		permission, err := space.CheckIsSpacePublicPermissionGranted(ctx, spaceId)
		if err != nil {
			slog.Error("error in space.CheckIsSpacePublicPermissionGranted, %w", err)
			ctx.AbortWithStatusJSON(500, model.NewInternalServerError())
			return
		}
		if !permission {
			ctx.AbortWithStatusJSON(403, NewNoPermissionError())
			return
		}
		ctx.Next()
	}
}
