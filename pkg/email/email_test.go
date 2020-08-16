package email

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("测试发送:", TestSend)
}
func TestSend(t *testing.T) {
	e := NewEmail(&SMTPInfo{
		Host:     "smtp.163.com",
		Port:     465,
		IsSSl:    true,
		UserName: "XXX@163.com",
		Password: "SSSSSS",
		From:     "XXX@163.com",
	})

	touser := []string{"XXX@qq.com"}
	err := e.SendMail(touser, "测试邮件", "wy测试邮件")
	if err != nil {
		fmt.Printf("邮件发送失败: %v\n", err)
	}
}
