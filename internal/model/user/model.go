package user

type User struct {
	// OpenId 用户的OpenID
	OpenId string `json:"openid" validate:"uuid" binding:"required"`
	// Name 用户的名字。
	Name string `json:"name"`
	// StudentId 用户的学号。
	StudentId string `json:"student_id"`
	// ClassName 用户的班级，顺便也表示了所属的学部院。
	ClassName string `json:"class_name"`
	// Group 所属组别。
	Group string `json:"group"`
	// Contact 联系方式
	Contact string `json:"contact"`
	// CreatedAt 创建时间，UNIX时间戳UTC时区以秒为单位
	CreatedAt int64 `json:"created_at"`
	// LastLoginAt 上次登录时间，UNIX时间戳UTC时区以秒为单位
	LastLoginAt int64 `json:"last_login_at"`
	// Meta 用户的元数据
	Meta string `json:"meta" validate:"json"`
}
