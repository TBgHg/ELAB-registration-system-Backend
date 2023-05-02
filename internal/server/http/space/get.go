package space

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

func NewSpaceNotFoundError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "space_not_found",
		Message: "空间不存在",
	}
}

func getSpaceById(ctx *gin.Context) {
	id := ctx.GetString("space_id")
	spaceResult, err := space.GetSpaceById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(404, NewSpaceNotFoundError())
			return
		}
		slog.Error("error in space.GetSpaceById, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, spaceResult)
}
