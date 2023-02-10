package mail

import (
	"fmt"
	"pAppServer/global"
	"time"

	"gopkg.in/gomail.v2"
)

type MailManager struct {
	MsgC       chan *gomail.Message
	SmtpSender gomail.SendCloser
}

func NewMailManager() *MailManager {
	mgr := new(MailManager)
	mgr.MsgC = make(chan *gomail.Message)
	return mgr
}

var MailMgr = NewMailManager()

func (m *MailManager) Send(mt MsgTytpe, to ...string) {
	var minfo *Msg
	if minfo = mt.Msg(); minfo == nil {
		return
	}
	mailConfig := global.Config.Mail
	msg := gomail.NewMessage()
	msg.SetHeader("From", msg.FormatAddress(mailConfig.UserName, mailConfig.Alias))
	msg.SetHeader("To", to...)
	msg.SetHeader("Subject", minfo.Subject)
	msg.SetBody("text/html", minfo.Body)
	MailMgr.MsgC <- msg
}

func Run() {
	mailConfig := global.Config.Mail
	d := gomail.NewDialer(mailConfig.Host, mailConfig.Port, mailConfig.UserName, mailConfig.Password)
	for {
		select {
		case msg := <-MailMgr.MsgC:
			fmt.Println("准备发送邮件")
			if MailMgr.SmtpSender == nil {
				fmt.Println("没有链接创建链接")
				if sender, err := d.Dial(); err != nil {
					panic(err)
				} else {
					MailMgr.SmtpSender = sender
				}
			}
			if err := gomail.Send(MailMgr.SmtpSender, msg); err != nil {
				fmt.Println("发送邮件失败", err.Error())
			} else {
				fmt.Println("发送邮件成功")
			}
		case <-time.After(time.Duration(mailConfig.CloseTime) * time.Second):
			fmt.Println("准备断开smtp链接")
			if MailMgr.SmtpSender != nil {
				fmt.Println("自动断开smtp链接")
				if err := MailMgr.SmtpSender.Close(); err != nil {
					panic(err)
				}
				MailMgr.SmtpSender = nil
			}
		}
	}

}
