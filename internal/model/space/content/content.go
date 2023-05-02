package content

import "gorm.io/gorm"

type Type string

const (
	Wiki    Type = "wiki"
	Thread  Type = "thread"
	Comment Type = "comment"
)

type OrderType string

const (
	TimeDesc OrderType = "time_desc"
	LikeDesc OrderType = "like_desc"
)

// Content 所谓内容，就是空间当中可被搜索的，以文本内容为主的内容。
//
// 目前包括：
//
//   - Wiki
//   - Thread
//   - Comment
type Content struct {
	gorm.Model `json:"-"` // 隐藏gorm.Model
	// SpaceId 空间的唯一标识符
	SpaceId string `json:"space_id" binding:"uuid"`
	// ContentId 内容的唯一标识符，内容标识符在一个空间当中是唯一的。
	ContentId string `json:"content_id" binding:"uuid"`
	// Type 内容的类型
	ContentType Type `json:"type"`
	// CurrentHistoryId 当前的历史版本的唯一标识符
	CurrentHistoryId string `json:"current_history_id" binding:"uuid"`
	// LastUpdateAt 最后更新时间，UNIX时间戳UTC时区以秒为单位
	LastUpdateAt int64 `json:"last_update_at"`
	// Meta 内容的元数据
	Meta string `json:"meta" binding:"json"`
}

// History 内容历史
type History struct {
	gorm.Model `json:"-"` // 隐藏gorm.Model
	// SpaceId 空间的唯一标识符
	SpaceId string `json:"space_id" binding:"uuid"`
	// ContentId 内容的唯一标识符
	ContentId string `json:"wiki_id" binding:"uuid"`
	// HistoryId 内容历史的唯一标识符
	HistoryId string `json:"history_id" binding:"uuid"`
	// OpenId 用户的OpenId
	OpenId string `json:"openid" binding:"uuid" gorm:"column:openid"`
	// Content 内容历史的内容
	Content string `json:"content"`
	// Time 内容历史的时间，UNIX时间戳UTC时区以秒为单位
	Time int64 `json:"time"`
	// Meta 内容历史的元数据
	Meta string `json:"meta" binding:"json"`
}

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
