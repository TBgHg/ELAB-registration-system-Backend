package member

import (
	"elab-backend/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	SendInvitation   Operation = "send_invitation"
	AcceptInvitation Operation = "accept_invitation"
	RejectInvitation Operation = "reject_invitation"
	RevokeInvitation Operation = "revoke_invitation"
)

type Invitation struct {
	gorm.Model `json:"-"`
	SpaceId    string                  `json:"space_id"`
	OpenId     string                  `json:"openid"`
	Status     InvitationOrApplyStatus `json:"status"`
}

func CheckIsInvitationExists(ctx *gin.Context, spaceId string, openId string) (bool, error) {
	svc := service.GetService()
	var count int64
	err := svc.MySQL.WithContext(ctx).Where(&Invitation{
		SpaceId: spaceId,
		OpenId:  openId,
		Status:  Pending,
	}).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// sendInvitation 给用户发送邀请，事先已经验证操作者与操作目标的身份，只需要确认邀请是否已经存在即可
func sendInvitation(ctx *gin.Context, request *OperateRequest) error {
	svc := service.GetService()
	invitationExists, err := CheckIsInvitationExists(ctx, request.SpaceId, request.OpenId)
	if err != nil {
		return err
	}
	if invitationExists {
		err = fmt.Errorf("invitation already exists")
		return err
	}
	invitation := Invitation{
		SpaceId: request.SpaceId,
		OpenId:  request.OpenId,
		Status:  Pending,
	}
	err = svc.MySQL.WithContext(ctx).Create(&invitation).Error
	return err
}

func acceptInvitation(ctx *gin.Context, request *OperateRequest) error {
	svc := service.GetService()
	invitationExists, err := CheckIsInvitationExists(ctx, request.SpaceId, request.OpenId)
	if err != nil {
		return err
	}
	if !invitationExists {
		err = fmt.Errorf("invitation not exists")
		return err
	}
	err = svc.MySQL.WithContext(ctx).Model(&Invitation{
		SpaceId: request.SpaceId,
		OpenId:  request.OpenId,
		Status:  Pending,
	}).Updates(&Invitation{
		Status: Accepted,
	}).Error
	if err != nil {
		return err
	}
	memberMeta := Meta{
		Position: NoPosition,
	}
	marshalledMeta, err := json.Marshal(memberMeta)
	if err != nil {
		return err
	}
	member := Member{
		SpaceId: request.SpaceId,
		OpenId:  request.OpenId,
		Meta:    string(marshalledMeta),
	}
	err = svc.MySQL.WithContext(ctx).Create(&member).Error
	return err
}

func rejectInvitation(ctx *gin.Context, request *OperateRequest) error {
	svc := service.GetService()
	invitationExists, err := CheckIsInvitationExists(ctx, request.SpaceId, request.OpenId)
	if err != nil {
		return err
	}
	if !invitationExists {
		err = fmt.Errorf("invitation not exists")
		return err
	}
	err = svc.MySQL.WithContext(ctx).Model(&Invitation{
		SpaceId: request.SpaceId,
		OpenId:  request.OpenId,
		Status:  Pending,
	}).Updates(&Invitation{
		Status: Rejected,
	}).Error
	return err
}

func revokeInvitation(ctx *gin.Context, request *OperateRequest) error {
	svc := service.GetService()
	invitationExists, err := CheckIsInvitationExists(ctx, request.SpaceId, request.OpenId)
	if err != nil {
		return err
	}
	if !invitationExists {
		err = fmt.Errorf("invitation not exists")
		return err
	}
	err = svc.MySQL.WithContext(ctx).Delete(&Invitation{
		SpaceId: request.SpaceId,
		OpenId:  request.OpenId,
		Status:  Pending,
	}).Error
	return err
}
