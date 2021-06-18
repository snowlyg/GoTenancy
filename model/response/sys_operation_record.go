package response

import "time"

type SysOperationRecord struct {
	TenancyResponse
	Ip           string        `json:"ip" `
	Method       string        `json:"method"`
	Path         string        `json:"path"`
	Status       int           `json:"status" `
	Latency      time.Duration `json:"latency" `
	Agent        string        `json:"agent" `
	ErrorMessage string        `json:"errorMessage" `
	Body         string        `json:"body"`
	Resp         string        `json:"resp"`
	UserID       uint          `json:"userId"`
	TenancyName  string        `json:"tenancyName"`
	UserName     string        `json:"userName"`
	NickName     string        `json:"nickName"`
}
