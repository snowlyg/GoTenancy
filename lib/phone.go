package lib

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成随机手机号码
func CreatePhoneNumber() string {
	prelist := []string{"130", "131", "132", "133", "134", "135", "136", "137", "138", "139", "147", "150", "151", "152", "153", "155", "156", "157", "158", "159", "186", "187", "188"}

	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	phoneNumber := prelist[rd.Intn(23)] + fmt.Sprintf("%08v", rd.Int31n(100000000))
	return phoneNumber
}
