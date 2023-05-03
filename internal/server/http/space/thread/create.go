package thread

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"elab-backend/internal/model/space/content/thread"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func createThread(ctx *gin.Context) {
	request := thread.CreateRequest{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	err = thread.Create(ctx, &request)
	if err != nil {
		slog.Error("error in thread.Create, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, space.OperationResponse{
		Ok: true,
	})
}
