package chat

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	Uid     int
	Conn    *websocket.Conn
	MsgC    chan []byte
	isClose bool
}

func NewChatClient(uid int, conn *websocket.Conn) *Client {
	return &Client{
		Uid:     uid,
		Conn:    conn,
		MsgC:    make(chan []byte),
		isClose: false, // 此标志防止顶号时channel被重复关闭
	}
}

func (c *Client) Listen() {
	go c.Receive()
	go c.Send()
}

func (c *Client) Close() bool {
	if c.isClose {
		return false
	}
	c.isClose = true
	close(c.MsgC)
	c.Conn.Close()
	return true
}

func (c *Client) Receive() {
	defer func() {
		ChatMgr.UngisterC <- c
	}()
	for {
		_, msg, err := c.Conn.ReadMessage()
		fmt.Println("收到客户端信息", string(msg))
		fmt.Println("错误信息", err)
		if err != nil {
			break
		}
	}
	fmt.Println("Receive协程停止了")
}

func (c *Client) Send() {
	for msg := range c.MsgC {
		c.Conn.WriteMessage(websocket.TextMessage, msg)
	}
	fmt.Println("Send协程停止了")
}
