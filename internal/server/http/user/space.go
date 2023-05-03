package user

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
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

func getUserSpaces(ctx *gin.Context) {
	spaces, err := space.GetUserSpaces(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, NewSpaceNotFoundError())
			return
		}
		slog.Error("user.GetUserSpaces failed", "error", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	if len(*spaces) == 0 {
		ctx.JSON(404, NewSpaceNotFoundError())
		return
	}
	ctx.JSON(200, space.UserSpaceListResponse{
		Spaces: *spaces,
	})
}
