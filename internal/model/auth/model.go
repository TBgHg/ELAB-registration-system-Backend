package auth

import "time"

const SessionUriScheme = "auth-session://"
const SessionExpiration = time.Minute * 10

// NewSessionRequestParams 使用State首先存储RedirectUri和code_verifier
// 用于 POST /auth/new
type NewSessionRequestParams struct {
	State         string `json:"state" binding:"required"`
	RedirectUri   string `json:"redirect_uri" binding:"required"`
	CodeChallenge string `json:"code_challenge" binding:"required"`
}

type Session struct {
	RedirectUri  string `json:"redirect_uri"`
	CodeVerifier string `json:"code_verifier"`
}

// NewSessionResponse 用于 POST /auth/new 的响应格式。
type NewSessionResponse struct {
	Ok bool `json:"ok"`
}

type RefreshRequestParams struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// CallbackRequestParams 用于OAuth2.0 CodeFlow 所规定的回调参数。
type CallbackRequestParams struct {
	Code  string `query:"code" binding:"required"`
	State string `query:"state" binding:"required"`
}

type CallbackResponseParams struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	State        string `json:"state"`
	CodeVerifier string `json:"code_verifier"`
}
