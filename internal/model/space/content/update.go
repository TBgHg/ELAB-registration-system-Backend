package content

import (
	"elab-backend/internal/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UpdateContentOptions struct {
	SpaceId     string
	ContentId   string
	ContentType Type
	Content     string
	ContentMeta interface{}
	HistoryMeta interface{}
}

func UpdateContent(ctx *gin.Context, options *UpdateContentOptions) error {
	svc := service.GetService()
	historyId := uuid.NewString()
	updateTime := time.Now().UTC().Unix()
	targetContent := Content{
		SpaceId:     options.SpaceId,
		ContentType: options.ContentType,
		ContentId:   options.ContentId,
	}
	err := svc.MySQL.First(&targetContent).Error
	if err != nil {
		return err
	}
	marshalledContentMeta, err := json.Marshal(options.ContentMeta)
	if err != nil {
		return err
	}
	targetContent.Meta = string(marshalledContentMeta)
	marshalledHistoryMeta, err := json.Marshal(options.HistoryMeta)
	if err != nil {
		return err
	}
	targetContent.CurrentHistoryId = historyId
	history := History{
		SpaceId:     options.SpaceId,
		ContentId:   options.ContentId,
		ContentType: options.ContentType,
		Content:     options.Content,
		HistoryId:   historyId,
		Time:        updateTime,
		Meta:        string(marshalledHistoryMeta),
	}
	err = svc.MySQL.Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Model(&targetContent).Updates(targetContent).Error
		if err != nil {
			return err
		}
		err = tx.WithContext(ctx).Create(&history).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
