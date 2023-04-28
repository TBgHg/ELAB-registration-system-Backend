package model

// ErrorResponse 错误响应
//
// 在这个项目里，不应该将正常的响应包装到一个结构体里，而是直接返回响应。
// 判断响应是否正常应该通过响应的状态码来判断。
type ErrorResponse struct {
	// Error 错误信息，例如 `unauthorized`
	Error string `json:"error"`
	// Message 错误信息
	Message string `json:"message"`
}
