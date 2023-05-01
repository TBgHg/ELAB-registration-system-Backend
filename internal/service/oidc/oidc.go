package oidc

import (
	"context"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/exp/slog"
	"golang.org/x/oauth2"
)

// Config OIDC 服务初始化选项
type Config struct {
	// Issuer OIDC 的发行者
	Issuer string
	// OAuthConfig OAuth 配置
	OAuthConfig *oauth2.Config
}

type Service struct {
	// Provider OIDC 提供者
	Provider *oidc.Provider
	// RemoteKeySet 远程密钥集，用于对 JWT 进行验证
	RemoteKeySet *oidc.RemoteKeySet
	// Verifier IdToken验证器，不过也能用于验证AccessToken
	Verifier *oidc.IDTokenVerifier
	// OAuthConfig OAuth 配置
	OAuthConfig *oauth2.Config
}

type ProviderClaims struct {
	// JwksUri OIDC 的公钥集地址
	JwksUri string `json:"jwks_uri"`
}

type UserInfoClaims struct {
	// PreferredUsername 用户名
	PreferredUsername string `json:"preferred_username"`
	// Picture 用户的头像
	Picture string `json:"picture"`
}

func (s *Service) RefreshToken(ctx context.Context, refreshToken string) (*oauth2.Token, error) {
	return s.OAuthConfig.TokenSource(ctx, &oauth2.Token{
		RefreshToken: refreshToken,
	}).Token()
}

// NewService 新服务
func NewService(config *Config) (*Service, error) {
	// 新建一个上下文
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, config.Issuer)
	if err != nil {
		err = fmt.Errorf("无法获取OIDC Provider: %w", err)
		slog.Error(err.Error())
		return nil, err
	}
	// github.com/coreos/go-oidc/v3 的OIDC使用策略如下：
	// 首先使用 oidc.NewProvider(ctx, config.Issuer) 获取一个 Provider
	// Provider仅提供了有限的字段，例如 Endpoint（通过provider.Endpoint()）
	// 对于Provider的 /.well-known/openid-configuration 其他字段
	// 会被解析到 ProviderClaims 结构体中
	// 此时应该使用provider.Claims(&providerClaims)将其解析到结构体中
	providerClaims := &ProviderClaims{}
	err = provider.Claims(providerClaims)
	if err != nil {
		err = fmt.Errorf("无法从OIDC 获取Claims: %w", err)
		slog.Error(err.Error())
		return nil, err
	}
	// RemoteKeySet 是一个远程密钥集，用于对 JWT 进行验证
	remoteKeySet := oidc.NewRemoteKeySet(ctx, providerClaims.JwksUri)
	// Verifier按道理讲应该是IDToken的验证工具
	// 不过神奇的发现，AccessToken和IDToken的结构体是一样的！
	verifier := provider.Verifier(&oidc.Config{ClientID: config.OAuthConfig.ClientID})
	service := &Service{
		Provider:     provider,
		RemoteKeySet: remoteKeySet,
		Verifier:     verifier,
		OAuthConfig:  config.OAuthConfig,
	}
	return service, nil
}
