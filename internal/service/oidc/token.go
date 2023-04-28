package oidc

import (
	"context"
	"golang.org/x/oauth2"
)

type TokenExchangeOptions struct {
	// Code 授权码
	Code string
	// RedirectUri 重定向地址
	RedirectUri string
	// State 状态
	State string
	// CodeVerifier PKCE所使用的code_verifier
	CodeVerifier string
}

func (s *Service) TokenExchange(ctx context.Context, options *TokenExchangeOptions) (*oauth2.Token, error) {
	return s.OAuthConfig.Exchange(ctx, options.Code,
		oauth2.SetAuthURLParam("code_verifier", options.CodeVerifier),
	)
}
