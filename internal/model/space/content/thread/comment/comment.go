package comment

import (
	"elab-backend/internal/model/space/content"
)

type Comment struct {
	CommentId    string         `json:"comment_id"`
	Content      string         `json:"content"`
	LastUpdateAt int64          `json:"last_updated_at"`
	Author       content.Author `json:"author"`
	LikeCounts   int64          `json:"likes"`
	Liked        bool           `json:"liked"`
}

type ContentMeta struct {
	ThreadId string
}

type HistoryMeta struct{}
