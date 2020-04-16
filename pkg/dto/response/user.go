package response

// UserAuthResponse 用户校验响应体
type UserAuthResponse struct {
	// 登录令牌
	Token string `json:"token"`
}
