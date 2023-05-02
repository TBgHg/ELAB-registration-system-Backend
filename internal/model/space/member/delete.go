package member

import (
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type DeleteMemberRequest struct {
	OpenId  string `form:"openid" binding:"required,uuid"`
	SpaceId string `form:"space_id" binding:"required,uuid"`
}

func DeleteMemberById(ctx *gin.Context, spaceId string, openid string) error {
	svc := service.GetService()
	return svc.MySQL.WithContext(ctx).Delete(&Member{
		SpaceId: spaceId,
		OpenId:  openid,
	}).Error
}
