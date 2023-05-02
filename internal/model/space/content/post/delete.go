package post

import (
	"elab-backend/internal/model/space/content"
	"github.com/gin-gonic/gin"
)

type DeleteRequest struct {
	SpaceId string
	PostId  string
}

func HardDelete(ctx *gin.Context, post *DeleteRequest, contentType content.Type) error {
	return content.HardDeleteContent(ctx, &content.HardDeleteContentOptions{
		SpaceId:     post.SpaceId,
		ContentId:   post.PostId,
		ContentType: contentType,
	})
}
