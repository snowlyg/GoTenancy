package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var SysConfigCategory = new(configCategory)

type configCategory struct{}

var configCategories = []model.SysConfigCategory{
	{Name: "商城配置", Key: "shop", Info: "商城配置", Sort: 0, Icon: "md-settings", Status: 1},
	{Name: "基础配置", Key: "base", Info: "基础配置", Sort: 0, Icon: "md-settings", Status: 1},
	{Name: "公众号配置", Key: "wechat", Info: "公众号配置", Sort: 0, Icon: "md-settings", Status: 1},
	{Name: "短信配置", Key: "message", Info: " 平台短信配置", Sort: 0, Icon: "chat-round", Status: 1},
	{Name: "商户基础配置", Key: "mer_base", Info: "商户基础配置", Sort: 0, Icon: "", Status: 1},
	{Name: "公众号支付配置", Key: "wechat_payment", Info: "公众号支付配置", Sort: 0, Icon: "", Status: 1},
	{Name: "阿里云配置", Key: "aliyun_oss", Info: "阿里云配置", Sort: 0, Icon: "", Status: 1},
	{Name: "七牛云配置", Key: "qiniuyun", Info: "七牛云配置", Sort: 0, Icon: "", Status: 1},
	{Name: "腾讯云配置", Key: "tengxun", Info: "腾讯云配置", Sort: 0, Icon: "", Status: 1},
	{Name: "小程序配置", Key: "smallapp", Info: "小程序配置", Sort: 0, Icon: "", Status: 1},
	{Name: "小程序支付配置", Key: "routine_pay", Info: "小程序支付配置", Sort: 0, Icon: "", Status: 1},
	{Name: "余额/充值设置", Key: "balance", Info: "余额/充值设置", Sort: 0, Icon: "", Status: 1},
	{Name: "分销配置", Key: "brokerage", Info: "分销配置", Sort: 0, Icon: "", Status: 1},
	{Name: "商户提货点配置", Key: "mer_take", Info: "商户提货点配置", Sort: 0, Icon: "", Status: 1},
	{Name: "支付宝支付配置", Key: "alipay", Info: "支付宝支付配置", Sort: 0, Icon: "", Status: 1},
}

func (m *configCategory) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.SysConfigCategory{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> config_categories 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&configCategories).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> config_categories 表初始数据成功!")
		return nil
	})
}
