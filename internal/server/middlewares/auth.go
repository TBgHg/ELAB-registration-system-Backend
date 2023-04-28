package middlewares

import (
	"elab-backend/internal/model"
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
	"strings"
)

func NewUnauthorizedErrorResponse() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "unauthorized",
		Message: "您尚未登录，请确认您的请求当中是否包含了Authorization Header。",
	}
}

func NewAuthorizationValueErrorResponse() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "authorization_value_error",
		Message: "您的Authorization Header格式不正确，请确认您的Authorization Header格式为Bearer <token>。",
	}
}

// OAuthMiddleware OAuth授权中间件
func OAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取Authorization Header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(403, NewUnauthorizedErrorResponse())
		}
		// 通过空格分割authHeader
		// authHeader格式为：Bearer <token>
		// 通过空格分割后，authHeader[0]为Bearer，authHeader[1]为<token>
		authHeaderSplit := strings.Split(authHeader, " ")
		if len(authHeaderSplit) != 2 {
			ctx.AbortWithStatusJSON(403, NewAuthorizationValueErrorResponse())
		}
		authType, accessToken := authHeaderSplit[0], authHeaderSplit[1]
		if authType != "Bearer" {
			ctx.AbortWithStatusJSON(403, NewAuthorizationValueErrorResponse())
		}
		svc := service.GetService()
		svc.Oidc.RemoteKeySet.VerifySignature(ctx, accessToken)
	}
}

func OAuthSelfValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
