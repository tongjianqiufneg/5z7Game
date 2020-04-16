package route

import (
	"5z7Game/http/handler"
	"github.com/gin-gonic/gin"
)

func Load(router *gin.Engine)  {
	router.GET("/", handler.IndexHandler)


}
