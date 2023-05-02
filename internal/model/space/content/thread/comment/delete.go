package comment

import (
	"elab-backend/internal/model/space/content"
	"github.com/gin-gonic/gin"
)

type DeleteRequest struct {
	SpaceId   string `form:"space_id" binding:"required,uuid"`
	CommentId string `form:"comment_id" binding:"required,uuid"`
}

func Delete(ctx *gin.Context, request *DeleteRequest) error {
	err := content.HardDeleteContent(
		ctx,
		&content.HardDeleteContentOptions{
			SpaceId:     request.SpaceId,
			ContentId:   request.CommentId,
			ContentType: content.Comment,
		})
	return err
}
