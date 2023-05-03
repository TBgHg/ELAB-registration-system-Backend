package content

import (
	"gorm.io/gorm"
)

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
	ContentType Type `json:"content_type"`
	// CurrentHistoryId 当前的历史版本的唯一标识符
	CurrentHistoryId string `json:"current_history_id" binding:"uuid"`
	// LastUpdateAt 最后更新时间，UNIX时间戳UTC时区以秒为单位
	LastUpdateAt int64 `json:"last_update_at"`
	// Meta 内容的元数据
	Meta string `json:"meta" binding:"json"`
}

type Author struct {
	OpenId string `json:"openid" binding:"uuid"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type Head struct {
	ContentId        string `json:"content_id" binding:"uuid"`
	CurrentHistoryId string `json:"current_history_id" binding:"uuid"`
	LastUpdateAt     int64  `json:"last_update_at"`
	Title            string `json:"title"`
	Summary          string `json:"summary"`
	Author           Author `json:"author"`
}

type SearchDatabaseResult struct {
	ContentId        string
	CurrentHistoryId string
	LastUpdateAt     int64
	OpenId           string `gorm:"column:openid"`
	Name             string
	Email            string
	Meta             string
}
