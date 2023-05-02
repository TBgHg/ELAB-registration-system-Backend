package comment

import (
	"elab-backend/internal/model/space/content"
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
	"time"
)

type LikeRequest struct {
	SpaceId   string `form:"space_id" binding:"required,uuid"`
	CommentId string `form:"comment_id" binding:"required,uuid"`
}

func CheckLike(ctx *gin.Context, request *LikeRequest) (bool, error) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return false, err
	}
	var count int64
	err = svc.MySQL.Model(&content.Like{
		SpaceId:     request.SpaceId,
		ContentId:   request.CommentId,
		OpenId:      token.Subject,
		ContentType: content.Comment,
	}).Count(&count).Error
	return count > 0, err
}

func Like(ctx *gin.Context, request *LikeRequest) error {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return err
	}
	currentTime := time.Now().UTC().Unix()
	return svc.MySQL.Create(&content.Like{
		SpaceId:     request.SpaceId,
		ContentId:   request.CommentId,
		OpenId:      token.Subject,
		ContentType: content.Comment,
		LikedAt:     currentTime,
	}).Error
}

func Dislike(ctx *gin.Context, request *LikeRequest) error {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return err
	}
	return svc.MySQL.Delete(&content.Like{
		SpaceId:     request.SpaceId,
		ContentId:   request.CommentId,
		OpenId:      token.Subject,
		ContentType: content.Comment,
	}).Error
}
