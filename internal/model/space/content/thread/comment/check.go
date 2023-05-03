package comment

import (
	"elab-backend/internal/model/space/content"
	"github.com/gin-gonic/gin"
)

type CheckExistsByIdOptions struct {
	SpaceId   string
	CommentId string
}

func CheckExistsById(ctx *gin.Context, options *CheckExistsByIdOptions) (bool, error) {
	target := content.CheckExistsByIdOptions{
		SpaceId:     options.SpaceId,
		ContentId:   options.CommentId,
		ContentType: content.Comment,
	}
	return content.CheckExistsById(ctx, &target)
}
