package tenant

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TenantBase struct {
	gorm.Model

	Name         string
	Mode         uint8
	Times        uint      //`剩余次数',
	ExpireTime   time.Time //'到期时间',
	CreationTime time.Time
	State        int8 //'数据状态0未审核，1审核未通过，2：审核通过，-1删除',
	IsTop        uint8
	Order        int64
	IsDel        int8    //
	Amount       float32 //'金钱',
	Price        float32 //'价格',
	Logo         string  //'头像',
	Tag          string
	AreaId       int64 //'所在区域代码',
	Province     string
	City         string
	County       string
	Addr         string
	Linkman      string
	Phone        string
	Lng          float64 // 经度
	Lat          float64 // 纬度
	Appid        uint
}

type Tenant struct {
	TenantBase

	FullName      string // 全称
	CertifyPics   string // 资质证明材料
	Desc          string
	Pics          string //'多图上传',
	Remark        string // 审核说明
	PermissionKey string
	TenantKey     string
	RabcUsers     []*RabcUser
}
