package manager

import (
	"pAppServer/manager/chat"
	"pAppServer/manager/mail"
)

func Run() {
	go chat.Run()
	go mail.Run()
}
