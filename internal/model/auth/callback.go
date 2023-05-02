package auth

// CallbackRequest 用于OAuth2.0 CodeFlow 所规定的回调参数。
type CallbackRequest struct {
	Code  string `query:"code" binding:"required"`
	State string `query:"state" binding:"required"`
}

type CallbackResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	State        string `json:"state"`
	CodeVerifier string `json:"code_verifier"`
}