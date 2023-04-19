package service

import (
	log "ELAB-registration-system-Backend/logger"
	"context"
	"fmt"
	"net/http"

	"ELAB-registration-system-Backend/internal/model"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
)

var provider *oidc.Provider

// OAuth2Init 使用init函数在包被导入时初始化Provider对象和Verifier对象
func OAuth2Init(issuer string) (err error) {
	// 创建一个Provider对象，传入auth服务器的发现文档地址
	provider, err = oidc.NewProvider(context.Background(), issuer)
	if err != nil {
		log.Logger.Error("NewService oidc.NewProvider failed err:" + err.Error())
		return
	}
	return
}

// VerifyToken 定义一个中间件函数，用来验证AccessToken和获取用户信息
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// todo: 从请求头中获取ClientID，方式待定
		clientID := c.GetHeader("ClientID")
		// 创建一个Verifier对象，传入ClientID和期望的令牌类型
		verifier := provider.Verifier(&oidc.Config{
			ClientID: clientID,
		})
		// 从请求头中获取Authorization字段
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// 如果没有Authorization字段，返回401错误
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			return
		}
		// 从Authorization字段中提取AccessToken
		var token string
		fmt.Sscanf(authHeader, "Bearer %s", &token)
		if token == "" {
			// 如果没有AccessToken，返回401错误
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Bearer token"})
			return
		}
		// 验证AccessToken的有效性和签名
		idToken, err := verifier.Verify(context.Background(), token)
		if err != nil {
			// 如果验证失败，返回401错误
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		// 解析AccessToken中的用户信息
		var claims model.TokenClaims
		err = idToken.Claims(&claims)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// 将用户信息存入上下文中，方便后续处理
		c.Set("user", claims)
		// 继续处理请求
		c.Next()
	}
}
