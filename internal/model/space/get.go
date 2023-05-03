package space

import (
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func GetSpaceById(ctx *gin.Context, id string) (*Space, error) {
	svc := service.GetService()
	var space Space
	err := svc.MySQL.WithContext(ctx).Model(&Space{
		SpaceId: id,
	}).First(&space).Error
	return &space, err
}
