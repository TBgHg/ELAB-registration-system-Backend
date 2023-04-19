package model

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	// todo: 根据实际情况添加需要的字段
	Email  string `json:"email"`
	OpenID string `json:"openid"`
	jwt.RegisteredClaims
}
