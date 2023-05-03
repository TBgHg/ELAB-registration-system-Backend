package auth

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/auth"
	"elab-backend/internal/model/user"
	"elab-backend/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/url"
)

func NewSessionNotFoundError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "session_not_found",
		Message: "无法找到您的Session，请确认您的Session是否正确或是否过期。",
	}
}

func SetCallbackResponseParams(query *url.Values, params auth.CallbackResponse) {
	query.Set("access_token", params.AccessToken)
	query.Set("refresh_token", params.RefreshToken)
	query.Set("state", params.State)
}

// callback 使用 code 换取 token
func callback(ctx *gin.Context) {
	params := auth.CallbackRequest{}
	err := ctx.ShouldBind(&params)
	if err != nil {
		slog.Error("ctx.ShouldBind failed", "error", err)
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	// 从Redis获取Session
	svc := service.GetService()
	session, err := auth.GetAuthSession(ctx, params.State)
	if err != nil {
		if errors.Is(err, auth.SessionNotFoundError{}) {
			slog.Error("auth.GetAuthSession failed", "error", err)
			ctx.JSON(404, NewSessionNotFoundError())
			return
		}
		slog.ErrorCtx(ctx, "redis.GetAuthSession failed", "error", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	// 使用code换取token
	token, err := svc.Oidc.TokenExchange(ctx, params.Code, session.CodeVerifier)
	if err != nil {
		slog.ErrorCtx(ctx, "oidc.TokenExchange failed", "error", err)
		slog.Error("code_verifier %v", session.CodeVerifier)
		ctx.JSON(500, model.NewInternalServerError())
	}
	ctx.Set("access_token", token.AccessToken)
	// 开始构建响应，需要access_token, refresh_token, state, code_verifier
	redirectUrl, err := url.Parse(session.RedirectUri)
	if err != nil {
		slog.ErrorCtx(ctx, "url.Parse failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	responseParams := auth.CallbackResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		State:        params.State,
	}
	query := redirectUrl.Query()
	SetCallbackResponseParams(&query, responseParams)
	redirectUrl.RawQuery = query.Encode()
	err = user.HandleAuthCallback(ctx)
	if err != nil {
		slog.ErrorCtx(ctx, "HandleAuthCallback failed", "error", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	// 再获取一次IdToken...
	idToken, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		slog.ErrorCtx(ctx, "oidc.GetToken failed", "error", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	// 更新用户最后一次登陆时间
	err = user.UpdateUserLastLogin(ctx, idToken.Subject)
	if err != nil {
		slog.ErrorCtx(ctx, "user.UpdateUserLastLogin failed", "error", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.Redirect(302, redirectUrl.String())
}
