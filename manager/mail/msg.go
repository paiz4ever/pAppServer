package mail

import (
	"fmt"
	"pAppServer/global"
)

type MsgTytpe int

type Msg struct {
	Subject string
	Body    string
}

const (
	RegisterMsg MsgTytpe = iota
)

func (m MsgTytpe) Msg() *Msg {

	switch m {

	case RegisterMsg:
		return &Msg{
			Subject: fmt.Sprintf("欢迎光临%v!", global.Config.App.Name),
			Body:    `<h1>尊敬的顾客：</h1><p>感谢您光临小店！</p><h2>本店的宗旨是：</h2><p>顾客至上</p><p>服务至上</p><p>品质至上</p>`,
		}
	default:
		return nil
	}
}
