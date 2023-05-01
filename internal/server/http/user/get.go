package user

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/user"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

func NewUserNotFoundError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "user_not_found",
		Message: "无法找到用户",
	}
}

func getUser(ctx *gin.Context) {
	userStruct, err := user.GetUser(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(404, NewUserNotFoundError())
			return
		}
		slog.Error("user.GetUser failed %w", err)
		ctx.JSON(403, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, userStruct)
}
