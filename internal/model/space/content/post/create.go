package post

import (
	"elab-backend/internal/model/space/content"
	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	SpaceId     string `form:"space_id" binding:"required,uuid"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	ContentType content.Type
}

func Create(ctx *gin.Context, post *CreateRequest) error {
	return content.Create(ctx, &content.CreateOptions{
		SpaceId:     post.SpaceId,
		Content:     post.Content,
		ContentType: post.ContentType,
		ContentMeta: ContentMeta{},
		HistoryMeta: HistoryMeta{
			Title:   post.Title,
			Summary: post.Content[:100],
		},
	})
}
