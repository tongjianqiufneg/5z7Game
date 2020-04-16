package handler

import (
	"5z7Game/http/response"
	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context)  {
	response.WrapContext(ctx).Success("hello,5z7")
}
