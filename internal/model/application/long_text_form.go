package application

import (
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LongTextForm 长文本表单
//
// 在Gin当中，会被用作Request
type LongTextForm struct {
	gorm.Model `json:"-"` // 不会被序列化
	// OpenId 用户OpenId
	OpenId string `json:"openid" binding:"required,uuid" gorm:"column:openid"`
	// 加入原因
	Reason string `json:"reason"`
	// 个人经历
	Experience string `json:"experience"`
	// 个人自我评价
	SelfEvaluation string `json:"self_evaluation"`
}

type UpdateLongTextFormResponse struct {
	Ok bool `json:"ok"`
}

func CheckLongTextFormExists(ctx *gin.Context) (bool, error) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return false, err
	}
	openid := token.Subject
	var counts int64
	err = svc.MySQL.WithContext(ctx).Model(
		&LongTextForm{
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

func UpdateLongTextForm(ctx *gin.Context, form *LongTextForm) error {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	openid := token.Subject
	if err != nil {
		return err
	}
	// 先检查是否存在
	exists, err := CheckLongTextFormExists(ctx)
	if err != nil {
		return err
	}
	form.OpenId = openid
	if !exists {
		err = svc.MySQL.WithContext(ctx).Create(&form).Error
		if err != nil {
			return err
		}
		return nil
	}
	err = svc.MySQL.WithContext(ctx).Model(&LongTextForm{
		OpenId: openid,
	}).Save(form).Error
	return err
}
