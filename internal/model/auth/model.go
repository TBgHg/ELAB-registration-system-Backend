package auth

// NewRequestParams 使用State首先存储RedirectUri
// 用于 POST /auth/new
type NewRequestParams struct {
	State       string `json:"state" binding:"required"`
	RedirectUri string `json:"redirect_uri" binding:"required"`
}

// NewResponse 用于 POST /auth/new 的响应格式。
type NewResponse struct {
	Ok bool `json:"ok"`
}

// CallbackRequestParams 用于OAuth2.0 CodeFlow 所规定的回调参数。
type CallbackRequestParams struct {
	Code  string `query:"code" binding:"required"`
	State string `query:"state" binding:"required"`
}
