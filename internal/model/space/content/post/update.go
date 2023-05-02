package post

import (
	"elab-backend/internal/model/space/content"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	Content     string `json:"content" binding:"required"`
	SpaceId     string `form:"space_id" binding:"required,uuid"`
	PostId      string `form:"post_id" binding:"required,uuid"`
	Title       string `json:"title" binding:"required"`
	ContentType content.Type
}

func Update(ctx *gin.Context, request *UpdateRequest) error {
	targetContent, targetHistory, err := content.GetContentAndHistoryById(ctx, &content.GetContentByIdOptions{
		SpaceId:     request.SpaceId,
		ContentId:   request.PostId,
		ContentType: request.ContentType,
	})
	if err != nil {
		return err
	}
	contentMeta := ContentMeta{}
	historyMeta := HistoryMeta{}
	err = json.Unmarshal([]byte(targetContent.Meta), &contentMeta)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(targetHistory.Meta), &historyMeta)
	if err != nil {
		return err
	}
	historyMeta.Title = request.Title
	historyMeta.Summary = request.Content[:100]
	return content.UpdateContent(ctx, &content.UpdateContentOptions{
		SpaceId:     request.SpaceId,
		ContentId:   request.PostId,
		ContentType: request.ContentType,
		Content:     request.Content,
		ContentMeta: contentMeta,
		HistoryMeta: historyMeta,
	})
}
