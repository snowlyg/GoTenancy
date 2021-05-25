package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/snowlyg/go-tenancy/g"
)

// SysMini 小程序
type SysMini struct {
	g.TENANCY_MODEL
	UUID      uuid.UUID `json:"uuid" gorm:"comment:UUID"`
	Name      string    `json:"name" form:"name" gorm:"type:varchar(50);column:name;comment:小程序名称"`
	AppID     string    `json:"appId" form:"appId" gorm:"type:varchar(100);column:app_id;comment:小程序appId"`
	AppSecret string    `json:"appSecret" form:"appSecret" gorm:"type:varchar(100);column:app_secret;comment:小程序appSecret"`
	Remark    string    `json:"remark" form:"remark" gorm:"type:varchar(254);column:remark;comment:小程序备注说明"`
}
