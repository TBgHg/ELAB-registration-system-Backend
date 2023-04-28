package user

import "gorm.io/gen"

type Querier interface {
	// Exists 判断用户是否存在
	//
	// EXISTS(SELECT * FROM @@table WHERE openid = @openId)
	Exists(openId string) (bool, error)
	// GetByOpenId 根据OpenId获取用户
	//
	// SELECT * FROM @@table WHERE openid = @openId
	GetByOpenId(openId string) (*gen.T, error)
}
