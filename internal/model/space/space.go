package space

import (
	"elab-backend/internal/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Query struct {
	Name     string `json:"name"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
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
	db = db.InnerJoins("members", "members.space_id = spaces.id")
	privateQuery := db.Model(&Space{Private: false})
	// 像这样的查询，可以作为一个SubQuery直接嵌入查询语句中
	userQuery := db.Model(&Member{OpenId: token.Subject})
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

// CreateSpace 创建一个新的空间，同时创建创建一个Member表，默认将自己添加进Member表里面
func CreateSpace(ctx *gin.Context, space *Space) error {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return err
	}
	space.Id = uuid.NewString()
	svc.MySQL.WithContext(ctx).Create(space)
	member := Member{
		SpaceId: space.Id,
		OpenId:  token.Subject,
	}
	memberMeta := MemberMeta{
		Position: Owner,
	}
	marshalledMeta, err := json.Marshal(memberMeta)
	if err != nil {
		return err
	}
	member.Meta = string(marshalledMeta)
	svc.MySQL.WithContext(ctx).Create(&member)
	return nil
}

func GetSpaceById(ctx *gin.Context, id string) (*Space, error) {
	svc := service.GetService()
	var space Space
	err := svc.MySQL.WithContext(ctx).Model(&Space{
		Id: id,
	}).First(&space).Error
	return &space, err
}

func CheckIsSpaceOwner(ctx *gin.Context, spaceId string) (bool, error) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return false, err
	}
	var member Member
	err = svc.MySQL.WithContext(ctx).Model(&Member{
		SpaceId: spaceId,
		OpenId:  token.Subject,
	}).First(&member).Error
	if err != nil {
		return false, err
	}
	var memberMeta MemberMeta
	err = json.Unmarshal([]byte(member.Meta), &memberMeta)
	if err != nil {
		return false, err
	}
	return memberMeta.Position == Owner, nil
}

func DeleteSpaceById(ctx *gin.Context, id string) error {
	svc := service.GetService()
	err := svc.MySQL.WithContext(ctx).Delete(&Space{
		Id: id,
	}).Error
	if err != nil {
		return err
	}
	err = svc.MySQL.WithContext(ctx).Delete(&Member{
		SpaceId: id,
	}).Error
	return err
}
