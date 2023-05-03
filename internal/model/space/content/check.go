package content

import (
	"elab-backend/internal/model/space/member"
	"elab-backend/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CheckExistsByIdOptions struct {
	SpaceId     string
	ContentId   string
	ContentType Type
}

func CheckExistsById(ctx *gin.Context, options *CheckExistsByIdOptions) (bool, error) {
	targetContent := Content{
		SpaceId:     options.SpaceId,
		ContentId:   options.ContentId,
		ContentType: options.ContentType,
	}
	svc := service.GetService()
	var count int64
	err := svc.MySQL.WithContext(ctx).Model(&Content{}).Where(&targetContent).Count(&count).Error
	return count > 0, err
}

type CheckHistoryExistsByIdOptions struct {
	SpaceId     string
	ContentId   string
	ContentType Type
	HistoryId   string
}

func CheckHistoryExistsById(ctx *gin.Context, options *CheckHistoryExistsByIdOptions) (bool, error) {
	targetHistory := History{
		SpaceId:     options.SpaceId,
		ContentId:   options.ContentId,
		HistoryId:   options.HistoryId,
		ContentType: options.ContentType,
	}
	svc := service.GetService()
	var count int64
	err := svc.MySQL.WithContext(ctx).Model(&History{}).Where(&targetHistory).Count(&count).Error
	return count > 0, err
}

func CheckOperatorPermission(ctx *gin.Context, contentType Type) (bool, error) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return false, nil
	}
	contentIdKey := ""
	switch contentType {
	case Thread:
		contentIdKey = "thread_id"
	case Comment:
		contentIdKey = "comment_id"
	case Wiki:
		contentIdKey = "wiki_id"
	default:
		return false, fmt.Errorf("invalid content type: %s", contentType)
	}
	spaceId := ctx.GetString("space_id")
	if spaceId == "" {
		return false, fmt.Errorf("space id not found in context")
	}
	contentId := ctx.GetString(contentIdKey)
	if contentId == "" {
		return false, fmt.Errorf("content id not found in context")
	}
	_, targetHistory, err := GetContentAndHistoryById(ctx, &GetContentByIdOptions{
		SpaceId:     spaceId,
		ContentId:   contentId,
		ContentType: contentType,
	})
	if err != nil {
		return false, err
	}
	if targetHistory.OpenId == token.Subject {
		return true, nil
	}
	position, err := member.GetSpacePosition(ctx, spaceId)
	if err != nil {
		return false, err
	}
	return position == member.Owner || position == member.Moderator, nil
}
