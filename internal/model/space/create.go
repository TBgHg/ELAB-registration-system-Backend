package space

import (
	"elab-backend/internal/model/space/member"
	"elab-backend/internal/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateSpace 创建一个新的空间，同时创建创建一个Member表，默认将自己添加进Member表里面
func CreateSpace(ctx *gin.Context, space *Space) error {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return err
	}
	space.SpaceId = uuid.NewString()
	svc.MySQL.WithContext(ctx).Create(space)
	selfMember := member.Member{
		SpaceId: space.SpaceId,
		OpenId:  token.Subject,
	}
	memberMeta := member.Meta{
		Position: member.Owner,
	}
	marshalledMeta, err := json.Marshal(memberMeta)
	if err != nil {
		return err
	}
	selfMember.Meta = string(marshalledMeta)
	svc.MySQL.WithContext(ctx).Create(&selfMember)
	return nil
}
