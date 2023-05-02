package post

import (
	"elab-backend/internal/model/space/content"
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type ContentResponse struct {
	Content string `json:"content"`
}

type GetContentByHistoryIdOptions struct {
	SpaceId     string
	PostId      string
	HistoryId   string
	ContentType content.Type
}

func GetContentByHistoryId(ctx *gin.Context, options *GetContentByHistoryIdOptions) (string, error) {
	svc := service.GetService()
	history := content.History{
		SpaceId:     options.SpaceId,
		ContentId:   options.PostId,
		ContentType: options.ContentType,
		HistoryId:   options.HistoryId,
	}
	err := svc.MySQL.WithContext(ctx).First(&history).Error
	if err != nil {
		return "", err
	}
	return history.ContentId, nil
}
