package member

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"elab-backend/internal/model/space/member"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func NewNotOwnerError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "not_owner",
		Message: "您不是空间的创建者，没有权限",
	}
}

func NewNotMemberError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "not_member",
		Message: "您不是空间的成员，没有权限",
	}
}

func NewNoPermissionError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "no_permission",
		Message: "您没有权限",
	}
}

func NewAlreadyMemberError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "already_member",
		Message: "该用户已经是空间成员",
	}
}

// operateMemberById 根据Id对用户的操作目前只有指定与撤除管理员权限。
// 该接口限定只有拥有者可以使用
func operateMemberById(ctx *gin.Context) {
	request := member.OperateByIdRequest{}
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
	if position != member.Owner {
		ctx.JSON(403, NewNotOwnerError())
		return
	}
	err = member.OperateMemberById(ctx, &request)
	if err != nil {
		slog.Error("error in space.OperateMemberById, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, space.OperationResponse{
		Ok: true,
	})
}

// operateMember 目前有下面几个操作：
//
// 1. 发出邀请 send_invitation
// 2. 接受邀请 accept_invitation
// 3. 拒绝邀请 reject_invitation
// 4. 撤销邀请 revoke_invitation
// 5. 加入申请 send_apply
// 6. 接受申请 accept_apply
// 7. 拒绝申请 reject_apply
// 8. 撤回申请 revoke_apply
func operateMember(ctx *gin.Context) {
	request := member.OperateRequest{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	// 对操作进行权限判定
	// 发出邀请、撤销邀请、接受申请、拒绝申请：自己为空间所有者或管理员，目标为非空间成员，且此时OpenId不得为空
	// 接受邀请、拒绝邀请、加入申请、撤回申请：自己为非空间成员
	switch request.Operation {
	case member.SendInvitation, member.RevokeInvitation, member.AcceptApply, member.RejectApply:
		if request.OpenId == "" {
			ctx.JSON(400, model.NewInvalidParamError())
			return
		}
		// 检测自己是否为空间成员以及是否为空间所有者或管理员
		isSelfMember, err := member.CheckIsMember(ctx, request.SpaceId)
		if err != nil {
			slog.Error("error in space.CheckIsMember, %w", err)
			ctx.JSON(500, model.NewInternalServerError())
			return
		}
		if !isSelfMember {
			ctx.JSON(403, NewNotMemberError())
			return
		}
		position, err := member.GetSpacePosition(ctx, request.SpaceId)
		if err != nil {
			slog.Error("error in space.GetSpacePosition, %w", err)
			ctx.JSON(500, model.NewInternalServerError())
			return
		}
		if position != member.Owner && position != member.Moderator {
			ctx.JSON(403, NewNoPermissionError())
			return
		}
		// 检测目标是否为非空间成员
		isMember, err := member.CheckIsMemberByOpenId(ctx, request.SpaceId, request.OpenId)
		if err != nil {
			slog.Error("error in space.CheckIsMemberByOpenId, %w", err)
			ctx.JSON(500, model.NewInternalServerError())
			return
		}
		if isMember {
			ctx.JSON(403, NewAlreadyMemberError())
			return
		}

	case member.AcceptInvitation, member.RejectInvitation, member.SendApply, member.RevokeApply:
		// 检测自己是否为空间成员
		isSelfMember, err := member.CheckIsMember(ctx, request.SpaceId)
		if err != nil {
			slog.Error("error in space.CheckIsMember, %w", err)
			ctx.JSON(500, model.NewInternalServerError())
			return
		}
		if isSelfMember {
			ctx.JSON(403, NewAlreadyMemberError())
			return
		}
	}
	err = member.OperateMember(ctx, &request)
	if err != nil {
		slog.Error("error in space.OperateMember, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, space.OperationResponse{
		Ok: true,
	})
}
