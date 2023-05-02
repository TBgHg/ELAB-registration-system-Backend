package middlewares

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"elab-backend/internal/model/space/content"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func NewSpaceIdRequiredError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "space_id_required",
		Message: "/space/:space_id is required",
	}
}

func NewWikiIdRequiredError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "wiki_id_required",
		Message: "/space/:space_id/wiki/:wiki_id is required",
	}
}

func NewThreadIdRequiredError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "thread_id_required",
		Message: "/space/:space_id/wiki/:wiki_id/thread/:thread_id is required",
	}
}

func NewCommentIdRequiredError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "comment_id_required",
		Message: "/space/:space_id/wiki/:wiki_id/thread/:thread_id/comment/:comment_id is required",
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

func SpacePrivatePermissionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 判断用户是否为Space成员
		spaceId := ctx.GetString("space_id")
		permission, err := space.CheckIsSpacePrivatePermissionGranted(ctx, spaceId)
		if err != nil {
			slog.Error("error in space.CheckIsSpacePrivatePermissionGranted, %w", err)
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

func WikiIdRequiredMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wikiId := ctx.GetString("wiki_id")
		if wikiId == "" {
			ctx.AbortWithStatusJSON(400, NewWikiIdRequiredError())
			return
		}
		ctx.Next()
	}
}

func ThreadIdRequiredMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		threadId := ctx.GetString("thread_id")
		if threadId == "" {
			ctx.AbortWithStatusJSON(400, NewThreadIdRequiredError())
			return
		}
		ctx.Next()
	}
}

func CommentIdRequiredMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		commentId := ctx.GetString("comment_id")
		if commentId == "" {
			ctx.AbortWithStatusJSON(400, NewCommentIdRequiredError())
			return
		}
		ctx.Next()
	}
}

// ContentOperatorMiddleware 检测用于操作的Content是否为本人的，或本人本身就是Space的管理员或者拥有者
func ContentOperatorMiddleware(contentType content.Type) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		permissionGranted, err := content.CheckOperatorPermission(ctx, contentType)
		if err != nil {
			slog.Error("error in content.CheckOperatorPermission, %w", err)
			ctx.AbortWithStatusJSON(500, model.NewInternalServerError())
			return
		}
		if !permissionGranted {
			ctx.AbortWithStatusJSON(403, NewNoPermissionError())
			return
		}
		ctx.Next()
	}
}
