package service

import (
	"github.com/snowlyg/go-tenancy/utils"
)

// EmailTest 发送邮件测试
func EmailTest() error {
	subject := "test"
	body := "test"
	return utils.EmailTest(subject, body)
}
