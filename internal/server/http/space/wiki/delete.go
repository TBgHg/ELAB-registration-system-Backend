package wiki

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"elab-backend/internal/model/space/content/wiki"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func NewNotFoundError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "wiki_not_found",
		Message: "未找到该 wiki",
	}
}

func deleteWikiById(ctx *gin.Context) {
	wikiId := ctx.GetString("wiki_id")
	spaceId := ctx.GetString("space_id")
	exists, err := wiki.CheckExistsById(ctx, spaceId, wikiId)
	if err != nil {
		slog.Error("error in wiki.CheckExistsById, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	if !exists {
		ctx.JSON(404, NewNotFoundError())
		return
	}
	err = wiki.HardDelete(ctx, spaceId, wikiId)
	if err != nil {
		slog.Error("error in wiki.HardDelete, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, space.OperationResponse{
		Ok: true,
	})
}
