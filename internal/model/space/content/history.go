package content

import "gorm.io/gorm"

// History 内容历史
type History struct {
	gorm.Model `json:"-"` // 隐藏gorm.Model
	// SpaceId 空间的唯一标识符
	SpaceId string `json:"space_id" binding:"uuid"`
	// ContentId 内容的唯一标识符
	ContentId string `json:"content_id" binding:"uuid"`
	// ContentType 内容的类型
	ContentType Type `json:"content_type" binding:"oneof=wiki thread comment"`
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
