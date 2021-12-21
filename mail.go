package u

import (
	"sync"
	"gopkg.in/gomail.v2"
)

type mail struct{
	mailHost string
	mailPort int
	fromUser string
	password string
	*gomail.Message
}

var _mail *mail
var _mailOnce sync.Once

// Mail 获取_mail对象
func Mail() *mail {
	_mailOnce.Do(func() {
		_mail = &mail{"smtp.qq.com", 465, "", "",gomail.NewMessage()}
	})
	return _mail
}

// SetMailServer 设置发送服务器
func (receiver *mail) SetMailServer(host string, port int) {
	receiver.mailHost = host
	receiver.mailPort = port
}

// SetAccount 设置发送账户
func (receiver *mail) SetAccount(fromUser string, password string) {
	receiver.fromUser = fromUser
	receiver.password = password
}

// SetSubject 设置主题
func (receiver *mail) SetSubject(subject string) {
	Mail().SetHeader("Subject", subject)
}

// SendMail 发送邮件
func (receiver *mail) SendMail(toUser ...string) error {
	Mail().SetHeader("From", receiver.fromUser )
	Mail().SetHeader("To", toUser... )
	dialer := gomail.NewDialer(receiver.mailHost, receiver.mailPort, receiver.fromUser, receiver.password)
	err := dialer.DialAndSend(Mail().Message)
	if err != nil {
		return err
	}
	return nil
}
