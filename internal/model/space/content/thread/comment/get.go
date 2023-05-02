package comment

import (
	"crypto/md5"
	"elab-backend/internal/model/space/content"
	"github.com/gin-gonic/gin"
)

type GetRequest struct {
	SpaceId  string `form:"space_id" binding:"required,uuid"`
	ThreadId string `form:"thread_id" binding:"required,uuid"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}

type GetResponse struct {
	Comments []Comment `json:"comments"`
}

type GetDbQueryResult struct {
	ContentId    string `gorm:"column:content_id"`
	Content      string `gorm:"column:content"`
	LastUpdateAt int64  `gorm:"column:last_update_at"`
	OpenId       string `gorm:"column:openid"`
	Name         string `gorm:"column:name"`
	Email        string `gorm:"column:email"`
	LikeCounts   int64  `gorm:"column:like_counts"`
	Liked        bool   `gorm:"column:liked"`
}

func Get(ctx *gin.Context, request *GetRequest) (*GetResponse, error) {
	db, err := BuildDatabaseQuery(ctx, request)
	if request.Page == 0 {
		request.Page = 1
	}
	if request.PageSize == 0 {
		request.PageSize = 10
	}
	offset := (request.Page - 1) * request.PageSize
	limit := request.PageSize
	if err != nil {
		return nil, err
	}
	var results []GetDbQueryResult
	err = db.Offset(offset).Limit(limit).Find(&results).Error
	if err != nil {
		return nil, err
	}
	var comments []Comment
	for _, result := range results {
		hash := md5.Sum([]byte(result.Email))
		avatar := "https://gravatar.loli.net/gravatar/" + string(hash[:])
		comments = append(comments, Comment{
			CommentId:    result.ContentId,
			Content:      result.Content,
			LastUpdateAt: result.LastUpdateAt,
			Author: content.Author{
				OpenId: result.OpenId,
				Name:   result.Name,
				Avatar: avatar,
			},
			LikeCounts: result.LikeCounts,
			Liked:      result.Liked,
		})
	}
	return &GetResponse{
		Comments: comments,
	}, nil
}
