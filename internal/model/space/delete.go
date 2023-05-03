package space

import (
	"elab-backend/internal/model/space/member"
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func DeleteSpaceById(ctx *gin.Context, id string) error {
	svc := service.GetService()
	err := svc.MySQL.WithContext(ctx).Delete(&Space{
		SpaceId: id,
	}).Error
	if err != nil {
		return err
	}
	err = svc.MySQL.WithContext(ctx).Delete(&member.Member{
		SpaceId: id,
	}).Error
	return err
}
