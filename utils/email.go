package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/snowlyg/go-tenancy/g"

	"github.com/jordan-wright/email"
)

// Email Email发送方法
func Email(subject string, body string) error {
	to := strings.Split(g.TENANCY_CONFIG.Email.To, ",")
	return send(to, subject, body)
}

// ErrorToEmail 给email中间件错误发送邮件到指定邮箱
func ErrorToEmail(subject string, body string) error {
	to := strings.Split(g.TENANCY_CONFIG.Email.To, ",")
	if to[len(to)-1] == "" { // 判断切片的最后一个元素是否为空,为空则移除
		to = to[:len(to)-1]
	}
	return send(to, subject, body)
}

// EmailTest Email测试
func EmailTest(subject string, body string) error {
	to := []string{g.TENANCY_CONFIG.Email.From}
	return send(to, subject, body)
}

// send Email发送方法
func send(to []string, subject string, body string) error {
	from := g.TENANCY_CONFIG.Email.From
	nickname := g.TENANCY_CONFIG.Email.Nickname
	secret := g.TENANCY_CONFIG.Email.Secret
	host := g.TENANCY_CONFIG.Email.Host
	port := g.TENANCY_CONFIG.Email.Port
	isSSL := g.TENANCY_CONFIG.Email.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}
