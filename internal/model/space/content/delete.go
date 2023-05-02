package content

import (
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HardDeleteContentOptions struct {
	SpaceId     string
	ContentId   string
	ContentType Type
}

func HardDeleteContent(ctx *gin.Context, options *HardDeleteContentOptions) error {
	// 删除所有对应的Content和History
	targetContent := Content{
		SpaceId:     options.SpaceId,
		ContentType: options.ContentType,
		ContentId:   options.ContentId,
	}
	targetHistory := History{
		SpaceId:     options.SpaceId,
		ContentType: options.ContentType,
		ContentId:   options.ContentId,
	}
	svc := service.GetService()
	err := svc.MySQL.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&targetContent).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&targetHistory).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
