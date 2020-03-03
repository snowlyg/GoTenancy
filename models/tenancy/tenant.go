package tenancy

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TenantBase struct {
	gorm.Model

	Name string
	Mode uint8
	//`剩余次数',
	Times uint
	//'到期时间',
	ExpireTime   *time.Time
	CreationTime *time.Time
	//'数据状态0未审核，1审核未通过，2：审核通过，-1删除',
	State int8
	IsTop uint8
	Order int64
	IsDel int8 //
	//'金钱',
	Amount float32
	//'价格',
	Price float32
	//'头像',
	Logo string
	Tag  string
	//'所在区域代码',
	AreaId   int64
	Province string
	City     string
	County   string
	Addr     string
	Linkman  string
	Phone    string
	// 经度
	Lng float64
	// 纬度
	Lat   float64
	Appid uint
}

type Tenant struct {
	TenantBase
	// 全称
	FullName string
	// 资质证明材料
	CertifyPics string
	Desc        string

	//'多图上传',
	Pics string
	// 审核说明
	Remark string

	PermissionKey string
	TenantKey     string
	RabcUsers     []*RabcUser
}
