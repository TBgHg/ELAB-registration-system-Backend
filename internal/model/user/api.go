package user

import (
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
	openid := ctx.Param("openid")
	userMatch := token.Subject == openid
	user := User{}
	db := svc.MySQL.WithContext(ctx).Model(&User{
		OpenId: openid,
	})
	if userMatch {
		db = db.Select("openid, name, `group`")
	}
	err = db.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUserLastLogin(ctx *gin.Context, openid string) error {
	svc := service.GetService()
	err := svc.MySQL.WithContext(ctx).Model(&User{}).Where(&User{
		OpenId: openid,
	}).Updates(&User{
		LastLoginAt: time.Now().UTC().Unix(),
	}).Error
	return err
}

func UpdateUser(ctx *gin.Context, openid string, updateBody User) error {
	svc := service.GetService()
	err := svc.MySQL.WithContext(ctx).Model(&User{}).Where(&User{
		OpenId: openid,
	}).Updates(&updateBody).Error
	return err
}
