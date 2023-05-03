package wiki

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"elab-backend/internal/model/space/content/wiki"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func updateWikiById(ctx *gin.Context) {
	request := wiki.UpdateRequest{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	exists, err := wiki.CheckExistsById(ctx, request.SpaceId, request.WikiId)
	if err != nil {
		slog.Error("error in wiki.CheckExistsById, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	if !exists {
		ctx.JSON(400, NewNotFoundError())
		return
	}
	err = wiki.Update(ctx, &request)
	if err != nil {
		slog.Error("error in wiki.Update, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, space.OperationResponse{
		Ok: true,
	})
}
