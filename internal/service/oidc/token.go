package oidc

import (
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"golang.org/x/oauth2"
)

type TokenExchangeOptions struct {
	// Code 授权码
	Code string
	// CodeVerifier PKCE所使用的code_verifier
	CodeVerifier string
}

func (s *Service) TokenExchange(ctx *gin.Context, code string, codeVerifier string) (*oauth2.Token, error) {
	return s.OAuthConfig.Exchange(ctx, code,
		oauth2.SetAuthURLParam("code_verifier", codeVerifier),
	)
}

func (s *Service) GetToken(ctx *gin.Context) (*oidc.IDToken, error) {
	accessToken, exists := ctx.Get("access_token")
	if !exists {
		err := fmt.Errorf("无法获取access_token")
		slog.Error(err.Error())
		return nil, err
	}
	token, err := s.Verifier.Verify(ctx, accessToken.(string))
	if err != nil {
		err = fmt.Errorf("无法验证access_token: %w", err)
		slog.Error(err.Error())
		return nil, err
	}
	return token, err
}
