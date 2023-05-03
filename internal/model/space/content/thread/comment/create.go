package comment

import (
	"elab-backend/internal/model/space/content"
	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	SpaceId  string `form:"space_id" binding:"required,uuid"`
	ThreadId string `form:"thread_id" binding:"required,uuid"`
	Content  string `json:"content" binding:"required"`
}

func Create(ctx *gin.Context, request *CreateRequest) error {
	contentMeta := ContentMeta{request.ThreadId}
	historyMeta := HistoryMeta{}
	return content.Create(ctx, &content.CreateOptions{
		SpaceId:     request.SpaceId,
		Content:     request.Content,
		ContentMeta: contentMeta,
		HistoryMeta: historyMeta,
		ContentType: content.Comment,
	})
}
