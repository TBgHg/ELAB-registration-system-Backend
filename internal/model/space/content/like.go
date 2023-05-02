package content

import "gorm.io/gorm"

// Like 内容点赞
type Like struct {
	gorm.Model `json:"-"` // 隐藏gorm.Model
	// SpaceId 空间的唯一标识符
	SpaceId string `json:"space_id" binding:"uuid"`
	// ContentId 内容的唯一标识符
	ContentId string `json:"content_id" binding:"uuid"`
	// OpenId 用户的OpenId
	OpenId string `json:"openid" binding:"uuid" gorm:"column:openid"`
	// LikedAt 点赞时间，UNIX时间戳UTC时区以秒为单位
	LikedAt int64 `json:"liked_at"`
}
