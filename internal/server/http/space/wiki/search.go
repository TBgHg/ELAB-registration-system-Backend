package wiki

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"elab-backend/internal/model/space/content"
	"elab-backend/internal/model/space/content/wiki"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func NewNoPermissionError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "no_permission",
		Message: "您没有权限",
	}
}

func searchWiki(ctx *gin.Context) {
	query := content.Query{}
	err := ctx.ShouldBind(&query)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	// 确认自己是否拥有相关权限
	isGranted, err := space.CheckIsSpacePublicPermissionGranted(ctx, query.SpaceId)
	if err != nil {
		slog.Error("error in space.CheckIsSpacePublicPermissionGranted, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	if !isGranted {
		ctx.JSON(403, NewNoPermissionError())
		return
	}
	heads, err := wiki.Search(ctx, &query)
	if err != nil {
		slog.Error("error in content.SearchContent, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, wiki.SearchResponse{
		Wikis: *heads,
	})
}
