package space

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func NewNotSpaceOwnerError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "not_space_owner",
		Message: "用户不是空间的创建者",
	}
}

// deleteSpaceById 根据ID删除空间
// 要求：用户必须是空间的创建者
func deleteSpaceById(ctx *gin.Context) {
	spaceId := ctx.GetString("space_id")
	isOwner, err := space.CheckIsSpaceOwner(ctx, spaceId)
	if err != nil {
		slog.Error("error in space.CheckIsSpaceOwner, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	if !isOwner {
		ctx.JSON(403, NewNotSpaceOwnerError())
		return
	}
	err = space.DeleteSpaceById(ctx, spaceId)
	if err != nil {
		slog.Error("error in space.DeleteSpaceById, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, space.OperationResponse{
		Ok: true,
	})
}
