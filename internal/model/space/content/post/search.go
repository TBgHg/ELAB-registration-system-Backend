package post

import (
	"elab-backend/internal/model/space/content"
	"github.com/gin-gonic/gin"
)

type SearchResponse struct {
	Posts []content.Head `json:"posts"`
}

func Search(ctx *gin.Context, query *content.Query, contentType content.Type) (*[]content.Head, error) {
	// 分两步处理
	// 1. 搜索Search
	// 2. 获取Meta进行解析
	// 初始化分页相关条件
	if query.PageSize <= 0 {
		query.PageSize = 10
	}
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.OrderBy == "" {
		query.OrderBy = content.TimeDesc
	}
	offset := (query.Page - 1) * query.PageSize
	size := query.PageSize
	db := BuildDatabaseQuery(ctx, query, contentType)
	// 至此，表格查询Column为
	// contents.*, histories.openid, users.email [_counts]
	// 由于属于自定义查表，因此需要重新定义一个结构体
	var dbQueryResult []content.SearchDatabaseResult
	err := db.Offset(offset).Limit(size).Scan(&dbQueryResult).Error
	if err != nil {
		return nil, err
	}
	result, err := ParseDatabaseSearchResult(&dbQueryResult)
	if err != nil {
		return nil, err
	}
	return result, nil
}
