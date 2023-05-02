package member

import (
	"elab-backend/internal/model/user"
	"elab-backend/internal/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type ListResponse struct {
	Members []Member `json:"members"`
}

func GetSpaceMemberList(ctx *gin.Context, spaceId string) (*[]Member, error) {
	svc := service.GetService()
	var members []Member
	err := svc.MySQL.WithContext(ctx).Model(&Member{SpaceId: spaceId}).Find(&members).Error
	if err != nil {
		return nil, err
	}
	userList := make([]user.User, len(members))
	var eg errgroup.Group
	for _, member := range members {
		member := member
		eg.Go(func() error {
			userQuery := user.User{
				OpenId: member.OpenId,
			}
			if err := svc.MySQL.WithContext(ctx).First(&userQuery).Error; err != nil {
				return err
			}
			userList = append(userList, userQuery)
			return nil
		})
	}
	if err = eg.Wait(); err != nil {
		return nil, err
	}
	return &members, nil
}

func CheckIsMember(ctx *gin.Context, spaceId string) (bool, error) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return false, err
	}
	return CheckIsMemberByOpenId(ctx, spaceId, token.Subject)
}

func CheckIsMemberByOpenId(ctx *gin.Context, spaceId string, openId string) (bool, error) {
	svc := service.GetService()
	var count int64
	err := svc.MySQL.WithContext(ctx).Model(&Member{
		SpaceId: spaceId,
		OpenId:  openId,
	}).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetSpacePosition(ctx *gin.Context, spaceId string) (Position, error) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return "", err
	}
	return GetSpacePositionByOpenId(ctx, spaceId, token.Subject)
}

func GetSpacePositionByOpenId(ctx *gin.Context, spaceId string, openId string) (Position, error) {
	svc := service.GetService()
	var member Member
	err := svc.MySQL.WithContext(ctx).Model(&Member{
		SpaceId: spaceId,
		OpenId:  openId,
	}).First(&member).Error
	if err != nil {
		return "", err
	}
	var memberMeta Meta
	err = json.Unmarshal([]byte(member.Meta), &memberMeta)
	if err != nil {
		return "", err
	}
	return memberMeta.Position, nil
}

// Member 空间成员
type Member struct {
	gorm.Model `json:"-"` // 隐藏gorm.Model
	// SpaceId 空间的唯一标识符
	SpaceId string `json:"space_id" binding:"uuid"`
	// OpenId 用户的OpenId
	OpenId string `json:"openid" binding:"uuid" gorm:"column:openid"`
	// Position 用户在空间中的职位
	//
	// 可能的值有：
	//  - owner
	//  - moderator
	//  - member
	Position string `json:"position"`
	// Meta 用户在空间中的元数据
	Meta string `json:"meta" binding:"json"`
}

type Position string

const (
	Owner      Position = "owner"
	Moderator  Position = "moderator"
	NoPosition Position = "none"
)

type Meta struct {
	// Position 用户在空间中的职位
	Position Position `json:"position"`
}
