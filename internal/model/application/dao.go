package application

import "gorm.io/gen"

type LongTextFormQuerier interface {
	// Exists 判断长文本表单是否存在
	//
	// EXISTS(SELECT * FROM @@table WHERE openid = @openId)
	Exists(openId string) (bool, error)
	// GetByOpenId 根据OpenId获取长文本表单
	//
	// SELECT * FROM @@table WHERE openid = @openId
	GetByOpenId(openId string) (*gen.T, error)
}
