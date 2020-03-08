package tenant

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media"
	"github.com/qor/media/oss"
)

type Tenant struct {
	gorm.Model

	Name      string
	FullName  string             // 全称
	Avatar    AvatarImageStorage //'头像',
	Province  string
	City      string
	County    string
	Addr      string
	Phone     string
	Lng       float64 // 经度
	Lat       float64 // 纬度
	RabcUsers []*RabcUser
}

type AvatarImageStorage struct{ oss.OSS }

func (AvatarImageStorage) GetSizes() map[string]*media.Size {
	return map[string]*media.Size{
		"small":  {Width: 50, Height: 50},
		"middle": {Width: 120, Height: 120},
		"big":    {Width: 320, Height: 320},
	}
}

func (tenant Tenant) AvatarImageURL() string {

	if &tenant.Avatar != nil && len(tenant.Avatar.URL("original")) > 0 {
		return tenant.Avatar.URL("original")
	}

	return "assets/images/avatars/3.jpg"
}
