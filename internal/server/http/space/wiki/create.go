package wiki

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"elab-backend/internal/model/space/content/wiki"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func createWiki(ctx *gin.Context) {
	request := wiki.CreateRequest{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	err = wiki.Create(ctx, &request)
	if err != nil {
		slog.Error("error in wiki.Create, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, space.OperationResponse{
		Ok: true,
	})
}
