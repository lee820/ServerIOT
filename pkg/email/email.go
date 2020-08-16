package email

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}

//SMTPInfo SMTP邮件信息
type SMTPInfo struct {
	Host     string
	Port     int
	IsSSl    bool
	UserName string
	Password string
	From     string
}

//NewEmail 新建Email结构体
func NewEmail(info *SMTPInfo) *Email {
	return &Email{SMTPInfo: info}
}

//SendMail 发送email
func (e *Email) SendMail(to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	//dialer smtp拨号器
	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSl}
	return dialer.DialAndSend(m)
}
