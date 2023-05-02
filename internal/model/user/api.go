package user

import (
	"elab-backend/internal/model/space"
	"elab-backend/internal/model/space/member"
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
	"time"
)

// CheckUserExists 检查用户是否存在，返回bool
func CheckUserExists(ctx *gin.Context) (bool, error) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return false, err
	}
	openid := token.Subject
	var counts int64
	err = svc.MySQL.WithContext(ctx).Model(&User{
		OpenId: openid,
	}).Count(&counts).Error
	if err != nil {
		return false, err
	}
	if counts == 0 {
		return false, nil
	}
	return true, nil
}

// InitUser 初始化用户
func InitUser(ctx *gin.Context) error {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return err
	}
	openid := token.Subject
	user := User{
		OpenId:       openid,
		IsElabMember: false,
		Meta:         "{}",
	}
	err = svc.MySQL.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func HandleAuthCallback(ctx *gin.Context) error {
	exists, err := CheckUserExists(ctx)
	if err != nil {
		return err
	}
	if !exists {
		err = InitUser(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetUser(ctx *gin.Context) (*User, error) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return nil, err
	}
	openid := ctx.GetString("openid")
	userMatch := token.Subject == openid
	user := User{}
	db := svc.MySQL.WithContext(ctx).Model(&User{
		OpenId: openid,
	})
	if userMatch {
		db = db.Select("openid, name, group")
	}
	err = db.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserSpaces(ctx *gin.Context) (*[]space.Space, error) {
	// 获取用户加入的空间，若access token和openid一致，显示全部内容，若不一致，仅显示private: false的空间
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return nil, err
	}
	openid := ctx.GetString("openid")
	userMatch := token.Subject == openid
	// 构建一个比较复杂的SQL查询：
	// Member当中，OpenId = openid
	// Space当中，ContentId = Member.SpaceId, Private = false
	db := svc.MySQL.Model(&space.Space{}).Joins(
		"JOIN member ON member.space_id = space.id").Where(&member.Member{OpenId: openid})
	if !userMatch {
		db = db.Where(&space.Space{Private: false})
	}
	var spaces []space.Space
	err = db.Find(&spaces).Error
	if err != nil {
		return nil, err
	}
	return &spaces, nil
}

func UpdateUserLastLogin(ctx *gin.Context, openid string) error {
	svc := service.GetService()
	err := svc.MySQL.WithContext(ctx).Model(&User{
		OpenId: openid,
	}).Updates(User{
		LastLoginAt: time.Now().UTC().Unix(),
	}).Error
	return err
}

func UpdateUser(ctx *gin.Context, openid string, updateBody User) error {
	svc := service.GetService()
	err := svc.MySQL.WithContext(ctx).Model(&User{
		OpenId: openid,
	}).Updates(updateBody).Error
	return err
}
