package space

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

// createSpace 根据传入的参数创建一个新的空间
func createSpace(ctx *gin.Context) {
	spaceRequest := space.Space{}
	err := ctx.ShouldBind(&spaceRequest)
	if err != nil {
		slog.Error("error in ctx.ShouldBind", "error", err)
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	if spaceRequest.Name == "" {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	err = space.CreateSpace(ctx, &spaceRequest)
	if err != nil {
		slog.Error("error in space.CreateSpace", "error", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, space.OperationResponse{
		Ok: true,
	})
}
