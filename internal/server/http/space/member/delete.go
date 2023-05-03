package member

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"elab-backend/internal/model/space/member"
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

// deleteMember 将用户从空间中移除
//
// 需要的权限：创建者或管理员，或者用户自己
func deleteMemberById(ctx *gin.Context) {
	svc := service.GetService()
	request := member.DeleteMemberRequest{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	isMember, err := member.CheckIsMember(ctx, request.SpaceId)
	if err != nil {
		slog.Error("error in space.CheckIsMember, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	if !isMember {
		ctx.JSON(403, NewNotMemberError())
		return
	}
	position, err := member.GetSpacePosition(ctx, request.SpaceId)
	if err != nil {
		slog.Error("error in space.GetSpacePosition, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	token, err := svc.Oidc.GetToken(ctx)
	isSelf := token.Subject == request.OpenId
	isOwnerOrModerator := position == "owner" || position == "moderator"
	if !isSelf && !isOwnerOrModerator {
		ctx.JSON(403, NewNoPermissionError())
		return
	}
	err = member.DeleteMemberById(ctx, request.SpaceId, request.OpenId)
	if err != nil {
		slog.Error("error in member.DeleteMemberById, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, space.OperationResponse{Ok: true})
}
