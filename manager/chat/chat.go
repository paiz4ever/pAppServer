package chat

import (
	"fmt"
	"pAppServer/global"
	"pAppServer/models/table"
)

type ChatManager struct {
	Clients    map[string]*Client
	RegisterC  chan *Client
	UngisterC  chan *Client
	BroadcastC chan []byte
}

func NewChatManager() *ChatManager {
	mgr := new(ChatManager)
	mgr.Clients = make(map[string]*Client)
	mgr.RegisterC = make(chan *Client)
	mgr.UngisterC = make(chan *Client)
	mgr.BroadcastC = make(chan []byte)
	return mgr
}

var ChatMgr = NewChatManager()

func (c *ChatManager) Send(msg []byte) {
	ChatMgr.BroadcastC <- msg
}

func Run() {
	for {
		select {
		case client := <-ChatMgr.RegisterC:
			if old, ok := ChatMgr.Clients[client.UUID]; ok {
				fmt.Println("顶号")
				old.Close()
			}
			client.Listen()
			ChatMgr.Clients[client.UUID] = client
			// TEMP
			var cuser table.CUser
			global.DB.Where("uuid = ?", client.UUID).First(&cuser)
			client.MsgC <- []byte(fmt.Sprint("欢迎你", cuser.NickName))
			fmt.Println("用户链接成功", client)
		case client := <-ChatMgr.UngisterC:
			if ok := client.Close(); ok {
				println("用户链接关闭")
				delete(ChatMgr.Clients, client.UUID)
			} else {
				println("用户被顶号时已经关闭")
			}
		case msg := <-ChatMgr.BroadcastC:
			for _, client := range ChatMgr.Clients {
				client.MsgC <- msg
			}
		}
	}
}
