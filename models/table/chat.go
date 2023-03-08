package table

import "gorm.io/gorm"

type ChatGroup struct {
	gorm.Model
	Name     string
	Notice   string
	MasterId int // 群主
	UserIds  string
}

type ChatMessage struct {
	gorm.Model
	FromId      int
	ToId        int
	MessageType int // 群 好友
	Content     int
	ContentType int
}
