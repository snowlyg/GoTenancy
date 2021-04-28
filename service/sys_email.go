package service

import (
	"github.com/snowlyg/go-tenancy/utils"
)

// EmailTest 发送邮件测试
func EmailTest() (err error) {
	subject := "test"
	body := "test"
	err = utils.EmailTest(subject, body)
	return err
}
