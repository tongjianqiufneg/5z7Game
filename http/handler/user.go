package handler

import (
	"5z7Game/http/response"
	"5z7Game/pkg/dto/request"
	"github.com/gin-gonic/gin"
)

// UserAuthHandler 用户登录
func UserAuthHandler(ctx *gin.Context)  {
	req := request.UserAuthRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {

	}
}

// UserRegisterHandler 用户注册
func UserRegisterHandler(ctx *gin.Context)  {
	// TODO
	response.WrapContext(ctx).Success(nil)
}