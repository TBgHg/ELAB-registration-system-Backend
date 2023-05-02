package thread

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"elab-backend/internal/model/space/content/thread"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func updateThreadById(ctx *gin.Context) {
	request := thread.UpdateRequest{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	exists, err := thread.CheckExistsById(ctx, request.SpaceId, request.ThreadId)
	if err != nil {
		slog.Error("error in thread.CheckExistsById, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	if !exists {
		ctx.JSON(400, NewNotFoundError())
		return
	}
	err = thread.Update(ctx, &request)
	if err != nil {
		slog.Error("error in thread.Update, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, space.OperationResponse{
		Ok: true,
	})
}
