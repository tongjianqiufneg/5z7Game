package request

// UserAuthRequest 用户验证请求
type UserAuthRequest struct {
	// 用户名
	Username string `json:"username"`

	// 密码
	Password string `json:"password"`
}
