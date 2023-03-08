package router

import (
	"net/http"
	"pAppServer/manager/chat"
	"pAppServer/models/response"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func hanleChatSocket(c *gin.Context) {
	uuid := c.GetString("uuid")
	if uuid == "" {
		response.Failed("uuid empty", c)
		return
	}
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	chat.ChatMgr.RegisterC <- chat.NewChatClient(uuid, conn)
}
