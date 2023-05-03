package member

import (
	"elab-backend/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ByIdOperation string

const (
	SetModerator    ByIdOperation = "set_moderator"
	RemoveModerator ByIdOperation = "remove_moderator"
)

type Operation string

type OperateByIdRequest struct {
	SpaceId   string        `form:"space_id" binding:"required,uuid"`
	OpenId    string        `form:"openid" binding:"required,uuid"`
	Operation ByIdOperation `json:"operation" binding:"required,oneof=set_moderator remove_moderator"`
}

type OperateRequest struct {
	SpaceId   string    `form:"space_id" binding:"required,uuid"`
	Operation Operation `json:"operation" binding:"required,oneof=send_invitation accept_invitation reject_invitation revoke_invitation join_apply accept_apply reject_apply revoke_apply"`
	OpenId    string    `json:"openid" binding:"uuid"`
}

type InvitationOrApplyStatus string

const (
	Pending  InvitationOrApplyStatus = "pending"
	Accepted InvitationOrApplyStatus = "accepted"
	Rejected InvitationOrApplyStatus = "rejected"
)

func OperateMemberById(ctx *gin.Context, request *OperateByIdRequest) error {
	svc := service.GetService()
	var member Member
	err := svc.MySQL.WithContext(ctx).Where(&Member{
		SpaceId: request.SpaceId,
		OpenId:  request.OpenId,
	}).First(&member).Error
	if err != nil {
		return err
	}
	memberMeta := Meta{}
	err = json.Unmarshal([]byte(member.Meta), &memberMeta)
	if err != nil {
		return err
	}
	switch request.Operation {
	case SetModerator:
		memberMeta.Position = Moderator
	case RemoveModerator:
		memberMeta.Position = NoPosition
	}
	marshalledMeta, err := json.Marshal(memberMeta)
	if err != nil {
		return err
	}
	member.Meta = string(marshalledMeta)
	err = svc.MySQL.WithContext(ctx).Save(&member).Error
	if err != nil {
		return err
	}
	return nil
}

func OperateMember(ctx *gin.Context, request *OperateRequest) error {
	switch request.Operation {
	case SendInvitation:
		return sendInvitation(ctx, request)
	case AcceptInvitation:
		return acceptInvitation(ctx, request)
	case RejectInvitation:
		return rejectInvitation(ctx, request)
	case RevokeInvitation:
		return revokeInvitation(ctx, request)
	case SendApply:
		return sendApply(ctx, request)
	case AcceptApply:
		return acceptApply(ctx, request)
	case RejectApply:
		return rejectApply(ctx, request)
	case RevokeApply:
		return revokeApply(ctx, request)
	default:
		return fmt.Errorf("invalid operation")
	}
}
