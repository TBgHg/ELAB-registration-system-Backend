package middlewares

import (
	"elab-backend/internal/model"
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"golang.org/x/oauth2"
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

func NewJwtInvalidErrorResponse() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "jwt_invalid",
		Message: "您的token无法通过验证，请确认您的token是否正确。",
	}
}

func NewGetOidcUserinfoFailedErrorResponse() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "get_oidc_userinfo_failed",
		Message: "无法获取用户信息，请确认您的token是否正确。",
	}
}

func NewUuidOidcSubMismatchErrorResponse() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "uuid_oidc_sub_mismatch",
		Message: "您的token与您的账号不匹配，请确认您的token是否正确。",
	}
}

// OAuthLoginRequiredMiddleware OAuth授权中间件
func OAuthLoginRequiredMiddleware() gin.HandlerFunc {
	svc := service.GetService()
	return func(ctx *gin.Context) {
		slog.DebugCtx(ctx, "OAuthLoginRequiredMiddleware")
		// 获取Authorization Header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(403, NewUnauthorizedErrorResponse())
			slog.DebugCtx(ctx, "aborted with status 403", "error", "unauthorized")
		}
		// 通过空格分割authHeader
		// authHeader格式为：Bearer <token>
		// 通过空格分割后，authHeader[0]为Bearer，authHeader[1]为<token>
		authHeaderSplit := strings.Split(authHeader, " ")
		if len(authHeaderSplit) != 2 {
			ctx.AbortWithStatusJSON(403, NewAuthorizationValueErrorResponse())
			slog.DebugCtx(ctx, "aborted with status 403", "error", "authorization_value_error")
		}
		authType, accessToken := authHeaderSplit[0], authHeaderSplit[1]
		if authType != "Bearer" {
			ctx.AbortWithStatusJSON(403, NewAuthorizationValueErrorResponse())
			slog.DebugCtx(ctx, "aborted with status 403", "error", "authorization_value_error")
		}
		// 验证token
		_, err := svc.Oidc.RemoteKeySet.VerifySignature(ctx, accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(403, NewJwtInvalidErrorResponse())
			slog.DebugCtx(ctx, "aborted with status 403", "error", "jwt_invalid", "detail", err)
		}
		ctx.Set("access_token", accessToken)
		ctx.Next()
	}
}

// OAuthSelfValidationMiddleware OAuth自我验证中间件
// 通过验证token中的uuid与oidc的sub是否一致来验证token是否属于当前用户
func OAuthSelfValidationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		slog.DebugCtx(ctx, "OAuthSelfValidationMiddleware")
		svc := service.GetService()
		uuid, exists := ctx.Get("uuid")
		if !exists {
			ctx.AbortWithStatusJSON(403, NewUuidOidcSubMismatchErrorResponse())
			slog.DebugCtx(ctx, "aborted with status 403", "error", "unauthorized")
			return
		}
		accessToken := ctx.MustGet("access_token").(string)
		userInfo, err := svc.Oidc.Provider.UserInfo(ctx,
			oauth2.StaticTokenSource(&oauth2.Token{
				AccessToken: accessToken,
			}))
		if err != nil {
			ctx.AbortWithStatusJSON(403, NewGetOidcUserinfoFailedErrorResponse())
			slog.DebugCtx(ctx, "aborted with status 403", "error", "get_userinfo_failed", "detail", err)
			return
		}
		if userInfo.Subject != uuid.(string) {
			ctx.AbortWithStatusJSON(403, NewJwtInvalidErrorResponse())
			slog.DebugCtx(ctx, "aborted with status 403", "error", "jwt_invalid", "detail", err)
		}
		ctx.Next()
	}
}
