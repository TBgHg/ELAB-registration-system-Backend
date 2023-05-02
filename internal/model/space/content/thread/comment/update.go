package comment

import (
	"elab-backend/internal/model/space/content"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	SpaceId   string `form:"space_id" binding:"required,uuid"`
	ThreadId  string `form:"thread_id" binding:"required,uuid"`
	CommentId string `form:"comment_id" binding:"required,uuid"`
	Content   string `json:"content" binding:"required"`
}

func Update(ctx *gin.Context, request *UpdateRequest) error {
	targetContent, targetHistory, err := content.GetContentAndHistoryById(ctx, &content.GetContentByIdOptions{
		SpaceId:     request.SpaceId,
		ContentId:   request.CommentId,
		ContentType: content.Comment,
	})
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
	err = content.UpdateContent(ctx, &content.UpdateContentOptions{
		SpaceId:     request.SpaceId,
		ContentId:   request.CommentId,
		ContentType: content.Comment,
		Content:     request.Content,
		ContentMeta: contentMeta,
		HistoryMeta: historyMeta,
	})
	return err
}
