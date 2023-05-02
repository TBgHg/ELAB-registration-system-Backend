package space

import "gorm.io/gorm"

// Space 空间
//
// 空间是OneELAB组织人员的最小单位。
type Space struct {
	gorm.Model `json:"-"` // 隐藏gorm.Model
	// SpaceId 空间的唯一标识符
	SpaceId string `json:"space_id" binding:"uuid"`
	// Name 空间的名称
	Name string `json:"name"`
	// Description 空间的描述
	Description string `json:"description"`
	// Private 空间是否为私有空间
	Private bool `json:"private"`
}

type ContentType string

const (
	ContentTypeWiki    ContentType = "wiki"
	ContentTypeThread  ContentType = "thread"
	ContentTypeComment ContentType = "comment"
)

type ContentOrderType string

const (
	ContentOrderByTimeDesc ContentOrderType = "time_desc"
	ContentOrderByLikeDesc ContentOrderType = "like_desc"
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
	ContentType ContentType `json:"type"`
	// CurrentHistoryId 当前的历史版本的唯一标识符
	CurrentHistoryId string `json:"current_history_id" binding:"uuid"`
	// LastUpdateAt 最后更新时间，UNIX时间戳UTC时区以秒为单位
	LastUpdateAt int64 `json:"last_update_at"`
	// Meta 内容的元数据
	Meta string `json:"meta" binding:"json"`
}

// ContentHistory 内容历史
type ContentHistory struct {
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

// ContentQuery 内容查询
type ContentQuery struct {
	// SpaceId 空间的唯一标识符
	SpaceId string `json:"space_id" binding:"uuid"`
	// Name 内容的名称
	Name string `json:"name"`
	// Type 内容的类型
	Type ContentType `json:"type"`
	// OrderBy 排序方式
	OrderBy ContentOrderType `json:"order_by"`
	// Limit 限制返回的内容数量
	Limit int32 `json:"limit"`
	// Offset 偏移量
	Offset int32 `json:"offset"`
}

// ContentLike 内容点赞
type ContentLike struct {
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

type OperationResponse struct {
	Ok bool `json:"ok"`
}
