package handler

import (
	"5z7Game/pkg/app"
	"github.com/ebar-go/ws"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebsocketHandler(ctx *gin.Context)  {
	conn , err := ws.GetUpgradeConnection(ctx.Writer, ctx.Request)
	if err != nil {
		http.NotFound(ctx.Writer, ctx.Request)
		return
	}

	client := ws.NewConnection(conn, func(ctx *ws.Context) string {
		// TODO 业务处理
		return ctx.GetMessage()
	})

	go client.Listen()

	app.Websocket().Register(client)

}
