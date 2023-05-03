package space

import (
	"elab-backend/internal/model/space/member"
	"elab-backend/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckIsSpacePublicPermissionGranted(ctx *gin.Context, spaceId string) (bool, error) {
	svc := service.GetService()
	// 判断能否有公开权限有两种方法：
	// 1. 判断空间是否不是私密空间
	// 2. 判断用户是否在空间中
	spaceQuery := Space{
		SpaceId: spaceId,
	}
	err := svc.MySQL.WithContext(ctx).First(&spaceQuery).Error
	if err != nil {
		return false, err
	}
	if !spaceQuery.Private {
		return true, nil
	}
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return false, err
	}
	targetMember := member.Member{
		SpaceId: spaceId,
		OpenId:  token.Subject,
	}
	err = svc.MySQL.WithContext(ctx).First(&targetMember).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CheckIsSpacePrivatePermissionGranted(ctx *gin.Context, spaceId string) (bool, error) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return false, err
	}
	targetMember := member.Member{
		SpaceId: spaceId,
		OpenId:  token.Subject,
	}
	var counts int64
	err = svc.MySQL.WithContext(ctx).Model(&member.Member{}).Where(&targetMember).Count(&counts).Error
	if err != nil {
		return false, err
	}
	return counts != 0, nil
}

// Space 空间
//
// 空间是OneELAB组织人员的最小单位。
type Space struct {
	gorm.Model `json:"-"` // 隐藏gorm.Model
	// SpaceId 空间的唯一标识符
	SpaceId string `json:"space_id"`
	// Name 空间的名称
	Name string `json:"name"`
	// Description 空间的描述
	Description string `json:"description"`
	// Private 空间是否为私有空间
	Private bool `json:"private"`
}

type OperationResponse struct {
	Ok bool `json:"ok"`
}
