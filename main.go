package main

import (
	"5z7Game/config"
	"5z7Game/http/route"
	"5z7Game/pkg/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func init()  {
	// 加载配置，错误则panic
	utils.SecurePanic(config.ReadFromFile("./app.yaml"))
}

func main() {
	router := gin.Default()

	route.Load(router)

	// go task.Run()

	_ = router.Run(":" + strconv.Itoa(config.Server().Port))
}
