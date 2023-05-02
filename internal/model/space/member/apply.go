package member

import (
	"elab-backend/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	SendApply   Operation = "send_apply"
	AcceptApply Operation = "accept_apply"
	RejectApply Operation = "reject_apply"
	RevokeApply Operation = "revoke_apply"
)

type Apply struct {
	gorm.Model `json:"-"`
	SpaceId    string                  `json:"space_id"`
	OpenId     string                  `json:"openid"`
	Status     InvitationOrApplyStatus `json:"status"`
}

func CheckIsApplyExists(ctx *gin.Context, spaceId string, openId string) (bool, error) {
	svc := service.GetService()
	var count int64
	err := svc.MySQL.WithContext(ctx).Where(&Apply{
		SpaceId: spaceId,
		OpenId:  openId,
		Status:  Pending,
	}).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func sendApply(ctx *gin.Context, request *OperateRequest) error {
	applyExists, err := CheckIsApplyExists(ctx, request.SpaceId, request.OpenId)
	if err != nil {
		return err
	}
	if applyExists {
		err = fmt.Errorf("apply already exists")
		return err
	}
	apply := Apply{
		SpaceId: request.SpaceId,
		OpenId:  request.OpenId,
		Status:  Pending,
	}
	svc := service.GetService()
	err = svc.MySQL.WithContext(ctx).Create(&apply).Error
	return err
}

func acceptApply(ctx *gin.Context, request *OperateRequest) error {
	svc := service.GetService()
	applyExists, err := CheckIsApplyExists(ctx, request.SpaceId, request.OpenId)
	if err != nil {
		return err
	}
	if !applyExists {
		err = fmt.Errorf("apply not exists")
		return err
	}
	err = svc.MySQL.WithContext(ctx).Model(&Apply{
		SpaceId: request.SpaceId,
		OpenId:  request.OpenId,
		Status:  Pending,
	}).Updates(&Apply{Status: Accepted}).Error
	memberMeta := Meta{
		Position: NoPosition,
	}
	marshalledMeta, err := json.Marshal(memberMeta)
	err = svc.MySQL.WithContext(ctx).Create(&Member{
		SpaceId: request.SpaceId,
		OpenId:  request.OpenId,
		Meta:    string(marshalledMeta),
	}).Error
	return err
}

func rejectApply(ctx *gin.Context, request *OperateRequest) error {
	svc := service.GetService()
	applyExists, err := CheckIsApplyExists(ctx, request.SpaceId, request.OpenId)
	if err != nil {
		return err
	}
	if !applyExists {
		err = fmt.Errorf("apply not exists")
		return err
	}
	err = svc.MySQL.WithContext(ctx).Model(&Apply{
		SpaceId: request.SpaceId,
		OpenId:  request.OpenId,
		Status:  Pending,
	}).Updates(&Apply{Status: Rejected}).Error
	return err
}

func revokeApply(ctx *gin.Context, request *OperateRequest) error {
	svc := service.GetService()
	applyExists, err := CheckIsApplyExists(ctx, request.SpaceId, request.OpenId)
	if err != nil {
		return err
	}
	if !applyExists {
		err = fmt.Errorf("apply not exists")
		return err
	}
	err = svc.MySQL.WithContext(ctx).Delete(&Apply{
		SpaceId: request.SpaceId,
		OpenId:  request.OpenId,
		Status:  Pending,
	}).Error
	return err
}
