package manager

import (
	"pAppServer/manager/chat"
	"pAppServer/manager/mail"
	"pAppServer/manager/wx"
)

func Run() {
	go chat.Run()
	go mail.Run()
	go wx.Run()
}
