package tenancy

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TenantBase struct {
	gorm.Model

	Name string `valid:"MaxSize(20)"`

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
	Logo string `valid:"MaxSize(255)"`
	Tag  string `valid:"MaxSize(255)"`
	//'所在区域代码',
	AreaId   int64
	Province string `valid:"MaxSize(20)"`
	City     string `valid:"MaxSize(20)"`
	County   string `valid:"MaxSize(20)"`
	Addr     string `valid:"MaxSize(20)"`
	Linkman  string `valid:"MaxSize(20)"`
	Phone    string `valid:"MaxSize(20)"`
	// 经度
	Lng float64
	// 纬度
	Lat   float64
	Appid uint
}

type Tenant struct {
	TenantBase
	// 全称
	FullName string `valid:"MaxSize(63)"`
	// 资质证明材料
	CertifyPics string `valid:"MaxSize(255)"`
	Desc        string

	//'多图上传',
	Pics string `valid:"MaxSize(255)"`
	// 审核说明
	Remark string `valid:"MaxSize(200)"`

	PermissionKey string `valid:"MaxSize(20)"`
	TenantKey     string `valid:"MaxSize(20)"`
	RabcUsers     []*RabcUser
}
