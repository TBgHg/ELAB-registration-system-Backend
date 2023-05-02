package space

import (
	"elab-backend/internal/model/space/member"
	"elab-backend/internal/service"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

// CreateSpace 创建一个新的空间，同时创建创建一个Member表，默认将自己添加进Member表里面
func CreateSpace(ctx *gin.Context, space *Space) error {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return err
	}
	space.SpaceId = uuid.NewString()
	svc.MySQL.WithContext(ctx).Create(space)
	selfMember := member.Member{
		SpaceId: space.SpaceId,
		OpenId:  token.Subject,
	}
	memberMeta := member.Meta{
		Position: member.Owner,
	}
	marshalledMeta, err := json.Marshal(memberMeta)
	if err != nil {
		return err
	}
	selfMember.Meta = string(marshalledMeta)
	svc.MySQL.WithContext(ctx).Create(&selfMember)
	return nil
}

func GetSpaceById(ctx *gin.Context, id string) (*Space, error) {
	svc := service.GetService()
	var space Space
	err := svc.MySQL.WithContext(ctx).Model(&Space{
		SpaceId: id,
	}).First(&space).Error
	return &space, err
}

func CheckIsSpacePublicPermissionGranted(ctx *gin.Context, spaceId string) (bool, error) {
	svc := service.GetService()
	// 判断能否有公开权限有两种方法：
	// 1. 判断空间是否不是私密空间
	// 2. 判断用户是否在空间中
	spaceQuery := Space{
		SpaceId: spaceId,
	}
	err := svc.MySQL.WithContext(ctx).First(&spaceQuery).Error
	if err != nil {
		return false, err
	}
	if !spaceQuery.Private {
		return true, nil
	}
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return false, err
	}
	targetMember := member.Member{
		SpaceId: spaceId,
		OpenId:  token.Subject,
	}
	err = svc.MySQL.WithContext(ctx).First(&targetMember).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func DeleteSpaceById(ctx *gin.Context, id string) error {
	svc := service.GetService()
	err := svc.MySQL.WithContext(ctx).Delete(&Space{
		SpaceId: id,
	}).Error
	if err != nil {
		return err
	}
	err = svc.MySQL.WithContext(ctx).Delete(&member.Member{
		SpaceId: id,
	}).Error
	return err
}

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

type OperationResponse struct {
	Ok bool `json:"ok"`
}
