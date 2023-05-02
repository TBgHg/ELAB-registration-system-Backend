package space

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func searchSpace(ctx *gin.Context) {
	query := space.Query{}
	err := ctx.ShouldBind(&query)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	spaces, err := space.SearchSpace(ctx, &query)
	if err != nil {
		slog.Error("error in space.SearchSpace, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, space.QueryResponse{
		Spaces: *spaces,
	})
}
