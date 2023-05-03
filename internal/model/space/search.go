package space

import (
	"elab-backend/internal/model/space/member"
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type Query struct {
	Name     string `form:"name"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}

type QueryResponse struct {
	Spaces []Space `json:"spaces"`
}

// SearchSpace 根据搜索条件搜索空间
// query为搜索条件，在自己已经加入的空间和公开空间中搜索
func SearchSpace(ctx *gin.Context, query *Query) (*[]Space, error) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return nil, err
	}
	var spaces []Space
	// 初始化分页相关条件
	if query.PageSize == 0 {
		query.PageSize = 10
	}
	if query.Page <= 0 {
		query.Page = 1
	}
	offset := (query.Page - 1) * query.PageSize
	size := query.PageSize
	db := svc.MySQL.WithContext(ctx)
	// 不做这个查询，将会合并spaces和members表，导致查询结果不正确
	db = db.Select("spaces.*")
	// 这是连接查询的重要部分
	db = db.InnerJoins("members", "members.space_id = spaces.space_id")
	privateQuery := db.Model(&Space{Private: false})
	// 像这样的查询，可以作为一个SubQuery直接嵌入查询语句中
	userQuery := db.Model(&member.Member{OpenId: token.Subject})
	if query.Name != "" {
		// 这里的?占位符可以嵌入子查询
		db = db.Where("spaces.name LIKE ? AND (? OR ?)", "%"+query.Name+"%", privateQuery, userQuery)
	} else {
		db = db.Where("(? OR ?)", privateQuery, userQuery)
	}
	db = db.Offset(offset).Limit(size)
	err = db.Find(&spaces).Error
	return &spaces, err
}
