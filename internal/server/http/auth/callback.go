package auth

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/auth"
	"elab-backend/internal/model/user"
	"elab-backend/internal/service"
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

func SetCallbackResponseParams(query *url.Values, params auth.CallbackResponseParams) {
	query.Set("access_token", params.AccessToken)
	query.Set("refresh_token", params.RefreshToken)
	query.Set("state", params.State)
	query.Set("code_verifier", params.CodeVerifier)
}

// callback 使用 code 换取 token
func callback(ctx *gin.Context) {
	params := auth.CallbackRequestParams{}
	err := ctx.ShouldBind(&params)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	// 从Redis获取Session
	svc := service.GetService()
	session, err := auth.GetAuthSession(ctx, params.State)
	if err != nil {
		slog.ErrorCtx(ctx, "redis.GetAuthSession failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	// 使用code换取token
	token, err := svc.Oidc.TokenExchange(ctx, params.Code, session.CodeVerifier)
	if err != nil {
		slog.ErrorCtx(ctx, "oidc.TokenExchange failed %w", err)
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
	responseParams := auth.CallbackResponseParams{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		State:        params.State,
		CodeVerifier: session.CodeVerifier,
	}
	query := redirectUrl.Query()
	SetCallbackResponseParams(&query, responseParams)
	redirectUrl.RawQuery = query.Encode()
	err = user.HandleAuthCallback(ctx)
	if err != nil {
		slog.ErrorCtx(ctx, "HandleAuthCallback failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	// 再获取一次IdToken...
	idToken, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		slog.ErrorCtx(ctx, "oidc.GetToken failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	// 更新用户最后一次登陆时间
	err = user.UpdateUserLastLogin(ctx, idToken.Subject)
	if err != nil {
		slog.ErrorCtx(ctx, "user.UpdateUserLastLogin failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.Redirect(302, redirectUrl.String())
}
