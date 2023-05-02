package content

import (
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type GetContentByIdOptions struct {
	SpaceId     string
	ContentId   string
	ContentType Type
}

type GetHistoryByIdOptions struct {
	SpaceId     string
	ContentId   string
	ContentType Type
	HistoryId   string
}

func GetContentById(ctx *gin.Context, options *GetContentByIdOptions) (*Content, error) {
	targetContent := Content{
		SpaceId:     options.SpaceId,
		ContentId:   options.ContentId,
		ContentType: options.ContentType,
	}
	svc := service.GetService()
	err := svc.MySQL.WithContext(ctx).First(&targetContent).Error
	return &targetContent, err
}

func GetHistoryById(ctx *gin.Context, options *GetHistoryByIdOptions) (*History, error) {
	targetHistory := History{
		SpaceId:     options.SpaceId,
		ContentId:   options.ContentId,
		ContentType: options.ContentType,
		HistoryId:   options.HistoryId,
	}
	svc := service.GetService()
	err := svc.MySQL.WithContext(ctx).First(&targetHistory).Error
	return &targetHistory, err
}

func GetContentAndHistoryById(ctx *gin.Context, options *GetContentByIdOptions) (*Content, *History, error) {
	targetContent, err := GetContentById(ctx, options)
	if err != nil {
		return nil, nil, err
	}
	targetHistory, err := GetHistoryById(ctx, &GetHistoryByIdOptions{
		SpaceId:     options.SpaceId,
		ContentId:   options.ContentId,
		ContentType: options.ContentType,
		HistoryId:   targetContent.CurrentHistoryId,
	})
	return targetContent, targetHistory, err
}
