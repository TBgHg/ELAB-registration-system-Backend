package auth

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/auth"
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func NewRefreshTokenFailedError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "refresh_token_failed",
		Message: "刷新token失败",
	}
}

// refresh 刷新token
func refresh(ctx *gin.Context) {
	params := auth.RefreshRequestParams{}
	err := ctx.ShouldBind(&params)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	svc := service.GetService()
	token, err := svc.Oidc.RefreshToken(ctx, params.RefreshToken)
	if err != nil {
		slog.Error("oidc.RefreshToken failed %w", err)
		ctx.JSON(403, NewRefreshTokenFailedError())
		return
	}
	ctx.JSON(200, auth.RefreshResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	})
}
