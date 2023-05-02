package wiki

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space/content/wiki"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func getWikiHistoryById(ctx *gin.Context) {
	request := wiki.GetContentByHistoryIdOptions{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	exists, err := wiki.CheckHistoryExistsById(ctx, request.SpaceId, request.WikiId, request.HistoryId)
	if err != nil {
		slog.Error("error in wiki.CheckHistoryExistsById, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	if !exists {
		ctx.JSON(400, NewNotFoundError())
		return
	}
	response, err := wiki.GetContentByHistoryId(ctx, &request)
	if err != nil {
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, wiki.ContentResponse{
		Content: response,
	})
}
