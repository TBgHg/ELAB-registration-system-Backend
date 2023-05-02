package content

import (
	"elab-backend/internal/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CreateOptions struct {
	SpaceId     string
	Content     string
	ContentType Type
	ContentMeta interface{}
	HistoryMeta interface{}
}

func Create(ctx *gin.Context, options *CreateOptions) error {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return err
	}
	contentId := uuid.NewString()
	historyId := uuid.NewString()
	marshalledContentMeta, err := json.Marshal(options.ContentMeta)
	if err != nil {
		return err
	}
	marshalledHistoryMeta, err := json.Marshal(options.HistoryMeta)
	if err != nil {
		return err
	}
	currentTime := time.Now().UTC().Unix()
	content := Content{
		SpaceId:          options.SpaceId,
		ContentId:        contentId,
		ContentType:      options.ContentType,
		CurrentHistoryId: historyId,
		LastUpdateAt:     currentTime,
		Meta:             string(marshalledContentMeta),
	}
	history := History{
		SpaceId:     options.SpaceId,
		ContentId:   contentId,
		ContentType: options.ContentType,
		HistoryId:   historyId,
		OpenId:      token.Subject,
		Content:     options.Content,
		Meta:        string(marshalledHistoryMeta),
		Time:        currentTime,
	}
	err = svc.MySQL.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&content).Error
		if err != nil {
			return err
		}
		err = tx.Create(&history).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
