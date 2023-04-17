package model

type TokenClaims struct {
	// todo: 根据实际情况添加需要的字段
	Email  string `json:"email"`
	OpenID string `json:"openid"`
}
